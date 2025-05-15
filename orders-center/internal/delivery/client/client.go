package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	url := c.baseURL + "/api/v1/order"

	data, err := json.Marshal(msg.Payload)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
