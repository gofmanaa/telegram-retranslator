package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
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

type Media struct {
	Url map[string]struct{}
}

func (m *Media) Add(url string) {
	m.Url[url] = struct{}{}
}

func Read(r []byte) {
	var data Posts
	media := Media{Url: make(map[string]struct{})}
	//err := json.Unmarshal(r, &data)
	buf := bytes.NewBuffer(r)
	err := json.NewDecoder(buf).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, post := range data {
		fmt.Printf("%s\n", post.Title.Rendered)
		//r := regexp.MustCompile(`<a[^>]+href=\"(.*?)\"[^>]*>(.*?)<\/a>`)
		//r := regexp.MustCompile(`(?:href=['"])([:\/.A-z?<_&\s=>0-9;-]+)`)
		//r := regexp.MustCompile(`(?:href=['"])([:\/.A-z?<_&\s=>0-9;-]+)`)
		//fmt.Printf("%s\n",r.FindAll([]byte(post.Content.Rendered), -1))

		//urls := r.FindAllString(post.Content.Rendered, -1)
		var re = regexp.MustCompile(`(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)`)
		var str = post.Content.Rendered

		for _, match := range re.FindAllString(str, -1) {
			media.Add(match)
		}

	}
		fmt.Println(media)
}
