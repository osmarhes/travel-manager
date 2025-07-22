# Build Stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY config/.env .env
CMD ["./main"]
