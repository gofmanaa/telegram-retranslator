package app

import (
	"context"
	"guthub.com/gofmanaa/telegram-bot/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/bot"
	"guthub.com/gofmanaa/telegram-bot/pkg/parser"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"time"
)

const WorkerCount = 5

func Run(conf *config.Config) {
	log := conf.Log
	defer log.Println("Stop application.")

	var ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rdb := redis_db.InitRedis(ctx, conf)

	var done chan struct{}
	ticker := time.NewTicker(time.Minute * 15)

	go func() {
		log.Printf("Start parser.SpaceGetto at %v\n", time.Now())
		collector := pool.StartDispatcher(ctx, WorkerCount) // start up worker pool
		for i, j := range parser.CreateJobs() {
			job := parser.SpaceGetto{
				Rdb:       rdb,
				InputData: j,
			}

			collector.Work <- pool.Work{Job: job, ID: i}
		}
		for t := range ticker.C {
			log.Printf("Start parser.SpaceGetto at %v\n", t)
			collector := pool.StartDispatcher(ctx, WorkerCount) // start up worker pool
			for i, j := range parser.CreateJobs() {
				job := parser.SpaceGetto{
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
