FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy

RUN go build -o /app/main ./cmd/avyaas/main.go

FROM alpine:3.18 AS deployer

WORKDIR /app

COPY --from=builder /app/main /app
COPY --from=builder /app/config.json /app

EXPOSE 9000

ENTRYPOINT ["/app/main"]
