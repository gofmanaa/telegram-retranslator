package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
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

type Set struct {
	sync.Mutex
	Set map[string]struct{}
}

func NewSet() *Set {
	return &Set{Set: make(map[string]struct{})}
}

func (s *Set) Add(data string) {
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
func CreateJobs(content []byte) []string {
	var inputData []string
	var post Posts

	buf := bytes.NewBuffer(content)
	err := json.NewDecoder(buf).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(post); i++ {
		inputData = append(inputData, post[i].Content.Rendered)
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
		if checkUrl(url) {
			err := sg.Rdb.Set(sg.Ctx, url, "", 0).Err()
			if err != nil {
				log.Fatal("Redis error: ", err)
			}
		}
	}
}

func checkUrl(url string) bool {
	if url == "" {
		return false
	}

	var re = regexp.MustCompile(`youtu*`)

	if re.MatchString(url) {
		return false
	}

	if existImage(url) != true {
		return false
	}

	return true
}

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
