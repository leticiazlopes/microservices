FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY microservices-proto/ ./microservices-proto/
COPY microservices/order/ ./microservices/order/

WORKDIR /app/microservices/order
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/microservices/order/order-service .
CMD ["./order-service"]