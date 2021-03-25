package parser

import (
	"bytes"
	"encoding/json"
	"fmt"

	"guthub.com/gofmanaa/telegram-bot/pkg/store"
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

func Scan(r []byte) *store.Media {
	var data Posts
	media := &store.Media{Url: make(map[string]struct{})}
	//err := json.Unmarshal(r, &data)
	buf := bytes.NewBuffer(r)
	err := json.NewDecoder(buf).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//wg := sync.WaitGroup{}
	fmt.Println("Posts count:", len(data))
	//resultCh := make(chan string, 500)
	postContents := []string{}

	for _, post := range data {
		postContents = append(postContents, post.Content.Rendered)
		// wg.Add(1)
		// go func(content string, result chan string) {
		// 	var re = regexp.MustCompile(`(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)`)

		// 	for _, match := range re.FindAllString(content, -1) {
		// 		result <- match
		// 	}
		// 	wg.Done()

		// }(post.Content.Rendered, ch)

	}
	// close(ch)
	//wg.Wait()
	// for val := range ch {
	// 	media.Add(val)
	// }

	return media
}
