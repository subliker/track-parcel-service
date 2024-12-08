FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/pkg

COPY internal/pkg/ .

WORKDIR /app/service/build

COPY internal/services/parcels_manager_service/go.mod .
COPY internal/services/parcels_manager_service/go.sum .
RUN go mod download

COPY internal/services/parcels_manager_service .
RUN go build -ldflags="-s -w" -o /app/parcels_manager_service ./cmd/pm


FROM alpine

WORKDIR /app/parcels_manager_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/parcels_manager_service /app/parcels_manager_service

CMD ["./parcels_manager_service"]
