package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ruandg/microservices-proto/golang/shipping"
	"github.com/ruandg/microservices/shipping/internal/adapters/db"
	"github.com/ruandg/microservices/shipping/internal/adapters/grpc"
	"github.com/ruandg/microservices/shipping/internal/application/core/api"
	realGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "3002"
	}

	dataSourceUrl := os.Getenv("DATA_SOURCE_URL")

	dbAdapter, err := db.NewAdapter(dataSourceUrl)
	if err != nil {
		log.Fatalf("Falha ao conectar no banco: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Falha ao abrir a porta %s: %v", port, err)
	}

	grpcServer := realGrpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, grpcAdapter)
	reflection.Register(grpcServer)

	go func() {
		log.Printf("Starting shipping service on port %s ...", port)
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Falha ao rodar o servidor gRPC: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Stopping shipping service...")
	grpcServer.GracefulStop()
}