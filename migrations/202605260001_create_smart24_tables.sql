-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    phone VARCHAR(30) NOT NULL,
    fullname VARCHAR(255),
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_code INTEGER UNIQUE,
    avatar TEXT
);

CREATE TABLE IF NOT EXISTS devices (
    id BIGSERIAL PRIMARY KEY,
    account VARCHAR(50) NOT NULL UNIQUE,
    code_device VARCHAR(100),
    user_code INTEGER,
    device_name VARCHAR(255),
    type INTEGER,
    bin VARCHAR(50),
    gruppa VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS device_set (
    id BIGSERIAL PRIMARY KEY,
    account VARCHAR(50) NOT NULL UNIQUE,
    user_wifi VARCHAR(255),
    wifi_pass VARCHAR(255),
    signal_wifi VARCHAR(50),
    status BOOLEAN DEFAULT FALSE,
    data_status TIMESTAMP,
    l INTEGER,
    n INTEGER,
    data_inkas TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tariffs (
    id BIGSERIAL PRIMARY KEY,
    tariff_id VARCHAR(50),
    name VARCHAR(255),
    sum NUMERIC(12,2) NOT NULL DEFAULT 0,
    type INTEGER
);

CREATE TABLE IF NOT EXISTS payments (
    id BIGSERIAL PRIMARY KEY,
    txn_id VARCHAR(255) NOT NULL UNIQUE,
    account VARCHAR(50) NOT NULL,
    sum NUMERIC(12,2) NOT NULL DEFAULT 0,
    result INTEGER DEFAULT 0,
    comment TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS coin (
    id BIGSERIAL PRIMARY KEY,
    account INTEGER NOT NULL,
    pay_coin INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS money (
    id BIGSERIAL PRIMARY KEY,
    account INTEGER NOT NULL,
    pay_money INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_devices_account ON devices(account);
CREATE INDEX IF NOT EXISTS idx_devices_user_code ON devices(user_code);
CREATE INDEX IF NOT EXISTS idx_device_set_account ON device_set(account);
CREATE INDEX IF NOT EXISTS idx_payments_account ON payments(account);
CREATE INDEX IF NOT EXISTS idx_coin_account ON coin(account);
CREATE INDEX IF NOT EXISTS idx_money_account ON money(account);

-- +goose Down
DROP TABLE IF EXISTS money;
DROP TABLE IF EXISTS coin;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS tariffs;
DROP TABLE IF EXISTS device_set;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS users;
