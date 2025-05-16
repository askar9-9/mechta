package config

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP       HTTPConfig
	Postgres   PostgresConfig
	OneC       OneCConfig
	Cron       CronConfig
	WorkerPool WorkerPoolConfig
}

type HTTPConfig struct {
	Addr       string        `env:"HTTP_ADDR" envDefault:":8080"`
	HealthAddr string        `env:"HTTP_HEALTH_ADDR" envDefault:":8081"`
	Timeout    time.Duration `env:"HTTP_TIMEOUT" envDefault:"30s"`
}

type PostgresConfig struct {
	Host         string        `env:"DB_HOST" envDefault:"localhost"`
	Port         int           `env:"DB_PORT" envDefault:"5432"`
	User         string        `env:"DB_USER" envDefault:"postgres"`
	Password     string        `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName       string        `env:"DB_NAME" envDefault:"orders"`
	SSLMode      string        `env:"DB_SSLMODE" envDefault:"disable"`
	MaxOpenConns int           `env:"DB_MAX_OPEN" envDefault:"10"`
	MaxIdleConns int           `env:"DB_MAX_IDLE" envDefault:"5"`
	ConnTimeout  time.Duration `env:"DB_CONN_TIMEOUT" envDefault:"5s"`
}

type OneCConfig struct {
	URL                   string        `env:"ONEC_URL" envDefault:"http://mock-1c:9900"`
	Timeout               time.Duration `env:"ONEC_TIMEOUT" envDefault:"5s"`
	Retries               int           `env:"ONEC_RETRIES" envDefault:"3"`
	MaxIdleConnections    int           `env:"ONEC_HTTP_MAX_IDLE" envDefault:"100"`
	MaxConnsPerHost       int           `env:"ONEC_HTTP_MAX_PER_HOST" envDefault:"10"`
	IdleConnTimeout       int           `env:"ONEC_HTTP_IDLE_TIMEOUT" envDefault:"90"`  // сек
	ResponseHeaderTimeout int           `env:"ONEC_HTTP_HEADER_TIMEOUT" envDefault:"5"` // сек
}

type WorkerPoolConfig struct {
	NumWorkers      int           `env:"WORKER_POOL_SIZE" envDefault:"10"`
	QueueSize       int           `env:"WORKER_POOL_QUEUE_SIZE" envDefault:"100"`
	DefaultTimeout  time.Duration `env:"WORKER_POOL_DEFAULT_TIMEOUT" envDefault:"5s"`
	MaxRetries      int           `env:"WORKER_POOL_MAX_RETRIES" envDefault:"3"`
	RetryBackoff    time.Duration `env:"WORKER_POOL_RETRY_BACKOFF" envDefault:"500ms"`
	ShutdownTimeout time.Duration `env:"WORKER_POOL_SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

type CronConfig struct {
	Interval time.Duration `env:"CRON_INTERVAL" envDefault:"60s"`
}

func New() (*Config, error) {
	cfg := new(Config)

	_ = godotenv.Load()

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	if cfg.Postgres.Host == "" || cfg.Postgres.DBName == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}

func (p PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode,
	)
}

func (c *Config) String() string {
	out, _ := json.MarshalIndent(c, "", "  ")
	return string(out)
}
