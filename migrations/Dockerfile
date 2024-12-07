FROM golang:alpine AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/goose /usr/local/bin/goose

WORKDIR /app

COPY . /app/migrations

CMD goose -dir /app/migrations/ postgres "host=$POSTGRES_HOST port=$POSTGRES_PORT user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable" up