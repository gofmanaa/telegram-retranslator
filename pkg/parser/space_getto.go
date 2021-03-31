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
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
}

// Unique data set
type Set struct {
	sync.Mutex
	Set map[string]struct{}
}

func NewSet() *Set {
	return &Set{Set: make(map[string]struct{})}
}

func (s *Set) Add(data string) {
	if checkUrl(data) != true {
		return
	}
	s.Lock()
	defer s.Unlock()
	s.Set[data] = struct{}{}
}

type SpaceGetto struct {
	Ctx       context.Context
	Rdb       *redis.Client
	InputData string
}

// create list of job
func CreateJobs() []string {
	var inputData []string
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

	for i := 0; i < len(posts); i++ {
		inputData = append(inputData, posts[i].Content.Rendered)
	}

	return inputData
}

// do job
func (sg SpaceGetto) DoWork() {
	imgs := NewSet()
	var re = regexp.MustCompile(`(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)`)
	for _, match := range re.FindAllString(sg.InputData, -1) {
		imgs.Add(match)
	}

	for url := range imgs.Set {
		if existImage(url) {
			redis_db.Save(sg.Ctx, sg.Rdb, url, 0)
		}
	}
}

func checkUrl(url string) bool {
	if url == "" {
		return false
	}

	//string shouldn't be youtube link (youtu.be|youtube.com)
	var re = regexp.MustCompile(`youtu*`)

	if re.MatchString(url) {
		return false
	}

	//url mast contain allow img type
	var regImg = regexp.MustCompile(`\.jpg$|\.png$|\.gif$|\.gifv$`)

	if !regImg.MatchString(url) {
		return false
	}

	return true
}

// Send Http Head request
func existImage(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Url [%s] Status [%d]", url, resp.StatusCode)
		return false
	}

	return true
}
