FROM golang:1.23-alpine AS builder

WORKDIR /app

# Устанавливаем swag для генерации Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Кэш go модулей
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Генерируем swagger docs на этапе build
RUN /go/bin/swag init -g cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o subscription-service ./cmd/main.go

# --- Release stage ---
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/subscription-service .
COPY --from=builder /app/docs ./docs
COPY .env .env

EXPOSE 8080

CMD ["./subscription-service"]