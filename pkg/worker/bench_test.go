package app

import (
	"guthub.com/gofmanaa/telegram-bot/cmd/app"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/pool"
	"guthub.com/gofmanaa/telegram-bot/pkg/worker/work"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(app.WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, srtJob := range work.CreateJobs(20) {
			job := work.TestJob{InputData: srtJob}
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, srtJob := range work.CreateJobs(20) {
			job := work.TestJob{InputData: srtJob}
			job.DoWork()
		}
	}
}
