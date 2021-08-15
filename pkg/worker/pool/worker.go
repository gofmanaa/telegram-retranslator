package pool

import (
	"context"
	"log"
)

type Jober interface {
	DoWork(ctx context.Context)
}

type Work struct {
	ID  int
	Job Jober
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel       chan Work
	End           chan bool
}

// start worker
func (w *Worker) Start(ctx context.Context) {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case job := <-w.Channel: // worker has received job
				job.Job.DoWork(ctx) // do work
			case <-w.End:
				return
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}
