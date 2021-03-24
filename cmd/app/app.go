package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"guthub.com/gofmanaa/telegram-bot/pkg/parser"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"io/ioutil"
)

const WORKER_COUNT = 5

var ctx = context.Background()

func Run() {
	fmt.Println("Start application.")
	defer fmt.Println("Stop application.")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool
	file, err := ioutil.ReadFile("tests/sg.json")
	if err != nil {
		fmt.Println(err)
	}
	for i, j := range parser.CreateJobs(file) {
		job := parser.SpaceGetto{
			Ctx:       ctx,
			Rdb:       rdb,
			InputData: j,
		}

		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
