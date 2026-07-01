FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY microservices-proto/ ./microservices-proto/
COPY microservices/payment/ ./microservices/payment/

WORKDIR /app/microservices/payment
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o payment-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/microservices/payment/payment-service .
CMD ["./payment-service"]