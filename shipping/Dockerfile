FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY microservices-proto/ ./microservices-proto/
COPY microservices/shipping/ ./microservices/shipping/

WORKDIR /app/microservices/shipping
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o shipping-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/microservices/shipping/shipping-service .
CMD ["./shipping-service"]