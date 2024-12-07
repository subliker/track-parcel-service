FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/pkg

COPY internal/pkg/ .

WORKDIR /app/service/build

COPY internal/services/account_service .
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/account_service ./cmd/acnt


FROM scratch

WORKDIR /app/account_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/account_service /app/account_service

CMD ["./account_service"]
