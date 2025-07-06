# syntax=docker/dockerfile:1

# Step 1: Build the Go binary
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Create a default .env if none exists
RUN if [ ! -f .env ] && [ -f .env.example ]; then cp .env.example .env; fi

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/api ./cmd/api

# Step 2: Create minimal final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates && \
    adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/api /app/api
# Copy environment file
COPY --from=builder /app/.env /app/.env

USER appuser

EXPOSE 4000

CMD ["/app/api"]