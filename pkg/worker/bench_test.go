package app

import (
	"context"
	"guthub.com/gofmanaa/telegram-bot/cmd/app"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/work"
	"testing"
)

var ctx = context.Background()

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(ctx, app.WORKER_COUNT, nil) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
