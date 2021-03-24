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
const JOB_COUNT = 100

var ctx = context.Background()

func Run() {
	fmt.Println("Start application.")
	defer fmt.Println("Stop application.")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sOmE_sEcUrE_pAsS", // no password set
		DB:       0,                  // use default DB
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
	//parser.Read(file)

	//	bot.Run()

	//err := rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key len: ", len(keys))
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}

}
