CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================
-- ENUM Types
-- ============================

-- Типы заказов (если фиксированные, например: pickup, delivery и т.д.)
CREATE TYPE order_type AS ENUM (
    'pickup',
    'delivery',
    'online'
);

-- Статусы заказа (примерные значения, уточни при необходимости)
CREATE TYPE order_status AS ENUM (
    'new',
    'processing',
    'confirmed',
    'shipped',
    'cancelled',
    'completed'
);

-- Типы оплаты
CREATE TYPE payment_type AS ENUM (
    'cash_at_shop',
    'cash_to_courier',
    'card',
    'card_online',
    'credit',
    'bonuses',
    'cashless',
    'prepayment'
);

-- ============================
-- Таблица: orders
-- ============================

CREATE TABLE orders (
    id VARCHAR(64) PRIMARY KEY,
    type order_type NOT NULL,
    status order_status NOT NULL,
    city VARCHAR(100) NOT NULL,
    subdivision VARCHAR(100),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    platform VARCHAR(64),
    general_id UUID NOT NULL UNIQUE,
    order_number VARCHAR(64) UNIQUE,
    executor VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_orders_general_id ON orders(general_id);
CREATE INDEX idx_orders_created_at ON orders(created_at);

-- ============================
-- Таблица: order_items
-- ============================

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    product_id VARCHAR(64) NOT NULL,
    external_id VARCHAR(64),
    status VARCHAR(64),
    base_price NUMERIC(10, 2) NOT NULL CHECK (base_price >= 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    earned_bonuses NUMERIC(10, 2) DEFAULT 0 CHECK (earned_bonuses >= 0),
    spent_bonuses NUMERIC(10, 2) DEFAULT 0 CHECK (spent_bonuses >= 0),
    gift BOOLEAN DEFAULT FALSE,
    owner_id VARCHAR(64),
    delivery_id VARCHAR(64),
    shop_assistant VARCHAR(100),
    warehouse VARCHAR(100),
    order_id UUID NOT NULL REFERENCES orders(general_id) ON DELETE CASCADE
);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);

-- ============================
-- Таблица: order_payments
-- ============================

CREATE TABLE order_payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL REFERENCES orders(general_id) ON DELETE CASCADE,
    type payment_type NOT NULL,
    sum NUMERIC(10, 2) NOT NULL CHECK (sum >= 0),
    payed BOOLEAN DEFAULT FALSE,
    info VARCHAR(255),
    contract_number VARCHAR(64),
    external_id VARCHAR(64)
);

CREATE INDEX idx_order_payments_order_id ON order_payments(order_id);

-- ============================
-- Таблица: credit_data
-- ============================

CREATE TABLE credit_data (
    id SERIAL PRIMARY KEY,
    order_payment_id UUID NOT NULL REFERENCES order_payments(id) ON DELETE CASCADE,
    bank VARCHAR(100),
    type VARCHAR(64),
    number_of_months SMALLINT CHECK (number_of_months > 0),
    pay_sum_per_month NUMERIC(10, 2),
    broker_id INT,
    iin VARCHAR(12) -- 12 символов для IIN в Казахстане
);

-- ============================
-- Таблица: card_payment_data
-- ============================

CREATE TABLE card_payment_data (
    id SERIAL PRIMARY KEY,
    order_payment_id UUID NOT NULL REFERENCES order_payments(id) ON DELETE CASCADE,
    provider VARCHAR(64),
    transaction_id VARCHAR(64)
);

-- ============================
-- Таблица: history
-- ============================

CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    type VARCHAR(64) NOT NULL,
    type_id INT,
    old_value BYTEA,
    value BYTEA,
    date TIMESTAMP NOT NULL DEFAULT now(),
    user_id VARCHAR(64),
    order_id VARCHAR(64) NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);

CREATE INDEX idx_history_order_id ON history(order_id);
CREATE INDEX idx_history_date ON history(date);


CREATE TABLE outbox_messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    aggregate_id UUID NOT NULL,
    aggregate_type VARCHAR(64) NOT NULL,
    event_type VARCHAR(64) NOT NULL,     
    payload JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    processed_at TIMESTAMP,
    retry_count INT NOT NULL DEFAULT 0,
    error TEXT
);

CREATE INDEX idx_outbox_unprocessed ON outbox_messages(processed_at) WHERE processed_at IS NULL;
