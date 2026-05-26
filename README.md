# Smart24 Backend

Go Gin + PostgreSQL + Goose migrations.

## Install goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Add Go bin to PATH:

```bash
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

## Setup

```bash
cp .env.example .env
nano .env
```

## Migrations

```bash
make migrate-up
make migrate-status
```

## Run API

```bash
go mod tidy
go run ./cmd/api
```

## Endpoints

```text
GET /health
GET /api/users
GET /api/devices
GET /api/device-set
GET /api/tariffs
GET /api/payments
GET /api/coin
GET /api/money
GET /api/devices/:account/full
GET /api/dashboard
GET /api/max
```

Pagination:

```text
/api/payments?limit=50&offset=0
```
