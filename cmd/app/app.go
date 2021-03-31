package app

import (
	"context"
	"fmt"
	"guthub.com/gofmanaa/telegram-bot/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/bot"
	"guthub.com/gofmanaa/telegram-bot/pkg/parser"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"time"
)

const WORKER_COUNT = 5
const DEADLINE = 50 * time.Second

var ctx = context.Background()

func Run(conf *config.Config) {
	fmt.Println("Start application.")
	defer fmt.Println("Stop application.")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rdb := redis_db.InitRedis(ctx, conf)

	var done chan struct{}
	ticker := time.NewTicker(time.Minute * 10)

	go func() {
		for t := range ticker.C {
			fmt.Printf("Start parser.SpaceGetto at %v\n", t)
			collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool
			for i, j := range parser.CreateJobs() {
				job := parser.SpaceGetto{
					Ctx:       ctx,
					Rdb:       rdb,
					InputData: j,
				}

				collector.Work <- pool.Work{Job: job, ID: i}
			}
			done <- struct{}{}
		}
	}()

	bot.Run(ctx, rdb, conf)

	<-done
}
