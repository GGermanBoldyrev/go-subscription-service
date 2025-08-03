FROM golang:1.23-alpine AS builder

WORKDIR /app

# Кэш go модулей
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o subscription-service ./cmd/main.go

# --- Release stage ---
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/subscription-service .
COPY .env .env

EXPOSE 8080

CMD ["./subscription-service"]