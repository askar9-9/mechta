package client

import (
	"context"
	"log/slog"
	"net/http"
	"orders-center/internal/application/config"
	"orders-center/internal/domain/outbox/entity"
)

type OneCClient struct {
	baseURL string
	client  *http.Client
}

func NewOneCClient(cfg *config.Config, client *http.Client) *OneCClient {
	return &OneCClient{
		baseURL: cfg.OneC.URL,
		client:  client,
	}
}

func (c *OneCClient) SendMessage(ctx context.Context, msg *entity.Outbox) error {
	slog.Info("SendMessage", "data", string(msg.Payload))
	return nil
}
