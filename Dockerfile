FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o smart24-api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/smart24-api .
COPY migrations ./migrations

EXPOSE 8081

CMD ["./smart24-api"]