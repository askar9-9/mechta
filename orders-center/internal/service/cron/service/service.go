package service

import (
	"context"
	"errors"
	"log/slog"
	"orders-center/internal/application/config"
	"runtime/debug"
	"sync"
	"time"
)

type WorkerPool struct {
	config     *config.WorkerPoolConfig
	taskQueue  chan Job
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         sync.WaitGroup
}

func NewWorkerPool(config *config.Config) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerPool{
		config:     &config.WorkerPool,
		taskQueue:  make(chan Job, config.WorkerPool.QueueSize),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.config.NumWorkers; i++ {
		w.wg.Add(1)
		go w.worker(i)
	}
	slog.Info("Worker pool started", "workers", w.config.NumWorkers)
}

func (w *WorkerPool) worker(id int) {
	defer w.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			stackTrace := debug.Stack()
			slog.Error("Worker panic recovered",
				"id", id,
				"error", r,
				"stack", string(stackTrace))

			// Restart the worker
			w.wg.Add(1)
			go w.worker(id)
		}
	}()

	slog.Info("Worker started", "id", id)
	for {
		select {
		case job, ok := <-w.taskQueue:
			if !ok {
				slog.Info("Worker shutting down (queue closed)", "id", id)
				return
			}

			w.processJob(job, id)
		case <-w.ctx.Done():
			slog.Info("Worker shutting down (context done)", "id", id)
			return
		}
	}
}

func (w *WorkerPool) processJob(job Job, workerID int) {
	jobID := job.ID()
	slog.Info("Processing job", "worker", workerID, "job", jobID)

	var err error
	for attempt := 0; attempt <= job.MaxRetries(); attempt++ {
		if attempt > 0 {
			slog.Info("Retrying job", "worker", workerID, "job", jobID, "attempt", attempt)
			time.Sleep(job.RetryDelay())
		}

		ctx, cancel := context.WithTimeout(context.Background(), job.Timeout())
		err = job.Process(ctx)

		if err == nil {
			if completeErr := job.OnComplete(ctx); completeErr != nil {
				slog.Error("Job completion failed",
					"err", completeErr,
					"worker", workerID,
					"job", jobID)
			}

			cancel()
			slog.Info("Job completed successfully", "worker", workerID, "job", jobID)
			return
		}

		errHandled := job.OnError(ctx, err)
		cancel()

		if errHandled == nil {
			slog.Info("Job error handled", "worker", workerID, "job", jobID)
			return
		}

		if attempt == job.MaxRetries() {
			slog.Error("Job failed after max retries",
				"err", err,
				"worker", workerID,
				"job", jobID,
				"attempts", attempt+1)
		}
	}
}

func (w *WorkerPool) Submit(ctx context.Context, job Job) error {
	select {
	case w.taskQueue <- job:
		slog.Info("Job submitted", "job", job.ID())
		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		return errors.New("task queue is full")
	}
}

func (w *WorkerPool) Stop() {
	slog.Info("Stopping worker pool...")

	w.cancelFunc()

	done := make(chan struct{})
	go func() {
		w.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		slog.Info("Worker pool stopped gracefully")
	case <-time.After(w.config.ShutdownTimeout):
		slog.Warn("Worker pool shutdown timed out")
	}

	close(w.taskQueue)
}

func (w *WorkerPool) QueueLength() int {
	return len(w.taskQueue)
}

func (w *WorkerPool) QueueCapacity() int {
	return cap(w.taskQueue)
}
