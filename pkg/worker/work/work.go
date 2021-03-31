package work

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type TestJob struct {
	InputData string
}

// create random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// create list of jobs
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(10))
	}
	return jobs
}

// mimics any type of job that can be run concurrently
func (tj TestJob) DoWork() {

	h := fnv.New32a()
	_, _ = h.Write([]byte(tj.InputData))
	time.Sleep(time.Second)
	fmt.Printf("Created hash [%d] from word [%s]\n", h.Sum32(), tj.InputData)
}
