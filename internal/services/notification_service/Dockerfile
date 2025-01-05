FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/pkg

COPY internal/pkg/ .

WORKDIR /app/service/build

COPY internal/services/notification_service/go.mod .
COPY internal/services/notification_service/go.sum .
RUN go mod download

COPY internal/services/notification_service .
RUN go build -ldflags="-s -w" -o /app/notification_service ./cmd/not


FROM alpine

WORKDIR /app/notification_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/notification_service /app/notification_service

CMD ["./notification_service"]
