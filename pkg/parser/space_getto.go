package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
)

type Posts []Post

type Post struct {
	ID    int `json:"id"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
}

// Unique data set
type Job struct {
	sync.Mutex
	ID      int
	Content string
}

type SpaceGetto struct {
	Rdb       *redis.Client
	InputData *Job
}

// create list of job
func CreateJobs() []*Job {
	var posts Posts
	resp, err := http.Get("https://www.spaceghetto.space/wp-json/wp/v2/posts")
	if err != nil {
		log.Panicln("Request error:", err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	buf := bytes.NewBuffer(data)
	if json.NewDecoder(buf).Decode(&posts) != nil {
		log.Fatal(err)
	}
	var jobs []*Job
	for _, post := range posts {
		job := &Job{}
		job.ID = post.ID
		job.Content = post.Content.Rendered
		jobs = append(jobs, job)
	}

	return jobs
}

type SGModel struct {
	mx     sync.Mutex
	ID     int
	Title  string
	Status int
	URLMap map[string]struct{}
}

func NewModel() *SGModel {
	return &SGModel{URLMap: make(map[string]struct{})}
}

func (m *SGModel) Save(ctx context.Context, client *redis.Client) {
	m.mx.Lock()
	has := redis_db.HasIdStatus(ctx, client, m.ID, "*")
	if has {
		m.mx.Unlock()
		return
	}
	redis_db.Set(ctx, client, m.ID, 0)
	for url := range m.URLMap {
		redis_db.SAdd(ctx, client, m.ID, url)
	}
	m.mx.Unlock()
}

// do job
func (sg SpaceGetto) DoWork(ctx context.Context) {
	model := NewModel()
	re := regexp.MustCompile(`(https:\/\/)([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)`)

	model.ID = sg.InputData.ID
	for _, match := range re.FindAllString(sg.InputData.Content, -1) {
		if checkUrl(match) {
			model.URLMap[match] = struct{}{}
		}
	}
	model.Save(ctx, sg.Rdb)
}

func checkUrl(url string) bool {
	if url == "" {
		return false
	}

	// string shouldn't be youtube link (youtu.be|youtube.com)
	var re = regexp.MustCompile(`youtu*`)

	if re.MatchString(url) {
		return false
	}

	// url mast contain allow img type
	var regImg = regexp.MustCompile(`\.jpg$|\.png$|\.gif$`)

	return regImg.MatchString(url)
}
