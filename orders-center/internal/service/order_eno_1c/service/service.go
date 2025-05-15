package service

import (
	"context"
	"fmt"
	"log/slog"
	"orders-center/internal/application/config"
	"orders-center/internal/delivery/client"
	outbox "orders-center/internal/domain/outbox/entity"
	"orders-center/internal/pkg/tx"
	eno "orders-center/internal/service/order_eno_1c/entity"
	"time"
)

type OrderEno1cService struct {
	syncInterval time.Duration
	workerCount  int
	oneCClient   *client.OneCClient
	outboxSvc    OutboxService
	cronSvc      CronService
	txManager    tx.TransactionManager
}

func NewOrderEno1cService(
	cfg *config.Config,
	client *client.OneCClient,
	outbox OutboxService,
	cron CronService,
	txManager tx.TransactionManager,
) *OrderEno1cService {
	return &OrderEno1cService{
		syncInterval: cfg.Cron.Interval,
		workerCount:  cfg.WorkerPool.NumWorkers,
		outboxSvc:    outbox,
		oneCClient:   client,
		cronSvc:      cron,
		txManager:    txManager,
	}
}

func (s *OrderEno1cService) Start(ctx context.Context) error {
	ticker := time.NewTicker(s.syncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := s.process(ctx); err != nil {
				slog.Error("process order_eno_1c err:", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *OrderEno1cService) Stop(ctx context.Context) error {
	ctx.Done()
	return nil
}

func (s *OrderEno1cService) Name() string {
	return "order_eno_1c"
}

func (s *OrderEno1cService) process(ctx context.Context) error {
	messages, err := s.getMessages(ctx)
	if err != nil {
		return fmt.Errorf("failed to get tasks: %w", err)
	}

	if len(messages) == 0 {
		slog.Info("no tasks to process")
		return nil
	}

	for _, msg := range messages {
		task := eno.NewTask(
			s.oneCClient,
			msg,
			func(ctx context.Context) error {
				return s.txManager.Do(ctx, func(ctx context.Context) error {
					if err := s.outboxSvc.UpdateSingle(ctx, msg); err != nil {
						return fmt.Errorf("failed to update message: %w", err)
					}
					return nil
				})
			})
		if err := s.cronSvc.Submit(ctx, task); err != nil {
			return fmt.Errorf("failed to submit task: %w", err)
		}
	}

	return nil
}

func (s *OrderEno1cService) getMessages(ctx context.Context) ([]*outbox.Outbox, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var messages []*outbox.Outbox

	// Use transaction manager to ensure atomicity
	err := s.txManager.Do(ctx, func(ctx context.Context) error {
		var err error
		messages, err = s.outboxSvc.FetchOutboxMessagesLimit(ctx, s.workerCount)
		if err != nil {
			return fmt.Errorf("failed to get message list: %w", err)
		}

		if len(messages) > 0 {
			// Set sync time for each task
			for _, msg := range messages {
				msg.SetSyncAt()
			}

			if err := s.outboxSvc.UpdateBatch(ctx, messages); err != nil {
				return fmt.Errorf("failed to update message list: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return messages, nil
}
