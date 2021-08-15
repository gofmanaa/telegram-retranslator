package app

import (
	"context"
	"guthub.com/gofmanaa/telegram-bot/cmd/telegram-bot/app"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/work"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	ctx := context.Background()
	collector := pool.StartDispatcher(ctx, app.WorkerCount) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, srtJob := range work.CreateJobs(20) {
			job := work.TestJob{InputData: srtJob}
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		for _, srtJob := range work.CreateJobs(20) {
			job := work.TestJob{InputData: srtJob}
			job.DoWork(ctx)
		}
	}
}
