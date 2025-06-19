# syntax=docker/dockerfile:1

# Step 1: Build the Go binary
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/api ./cmd/api

# Step 2: Create minimal final image
FROM alpine:latest

RUN adduser -D -g '' appuser

COPY --from=builder /bin/api /bin/api

USER appuser

EXPOSE 4000

CMD ["/bin/api"]