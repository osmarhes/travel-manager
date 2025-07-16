# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .
CMD ["./main"]
