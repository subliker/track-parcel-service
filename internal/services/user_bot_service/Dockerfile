FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/pkg

COPY internal/pkg/ .

WORKDIR /app/service/build

COPY internal/services/user_bot_service/go.mod .
COPY internal/services/user_bot_service/go.sum .
RUN go mod download

COPY internal/services/user_bot_service .
RUN go build -ldflags="-s -w" -o /app/user_bot_service ./cmd/ubot


FROM scratch

WORKDIR /app/user_bot_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/user_bot_service /app/user_bot_service

CMD ["./user_bot_service"]
