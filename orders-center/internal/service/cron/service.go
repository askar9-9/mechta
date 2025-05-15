package cron

import (
	"context"
	"log/slog"
	"orders-center/internal/application/config"
	"sync"
	"time"
)

type WorkerPool struct {
	workers    int
	taskQueue  chan Job
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         sync.WaitGroup
}

func NewWorkerPool(cfg *config.Config) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workers:    cfg.WorkerPool.NumWorkers,
		taskQueue:  make(chan Job, cfg.WorkerPool.NumWorkers),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go func(workerID int) {
			defer w.wg.Done()
			for {
				select {
				case job, ok := <-w.taskQueue:
					if !ok {
						return
					}

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					if err := job.Process(ctx); err != nil {
						slog.Error("job process failed", "err", err, "worker", workerID)
					}
					if err := job.OnComplete(ctx); err != nil {
						slog.Error("job completion failed", "err", err, "worker", workerID)
					}
				case <-w.ctx.Done():
					return
				}
			}
		}(i)
	}
}

func (w *WorkerPool) Submit(ctx context.Context, job Job) error {
	select {
	case w.taskQueue <- job:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (w *WorkerPool) Stop() {
	w.cancelFunc()
	w.wg.Wait()
	close(w.taskQueue)
}
