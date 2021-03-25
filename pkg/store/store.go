package store

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
)

type MediaUrl struct {
	Url string
}

type Media struct {
	sync.Mutex
	Url map[string]struct{}
}

func (m *Media) Add(url string) {
	m.Lock()
	if filter(url) {
		m.Url[url] = struct{}{}
	}
	m.Unlock()
}

func filter(url string) bool {
	if !youtube(url) {
		//&& existsUrl(url)

		return true
	}

	return false
}

func youtube(url string) bool {
	var re = regexp.MustCompile(`youtu.?`)
	return re.Match([]byte(url))
}

func existsUrl(url string) bool {
	res, err := http.Head(url)

	if err != nil {
		return false
	}
	if res.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func (m *Media) Save(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error write file: %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for line := range m.Url {
		_, err = fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
		//w.WriteString(line)
	}

	return w.Flush()
}
