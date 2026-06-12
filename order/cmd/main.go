package main

import (
	"log"

	"github.com/ruandg/microservices/order/config"
	"github.com/ruandg/microservices/order/internal/adapters/db"
	"github.com/ruandg/microservices/order/internal/adapters/grpc"
	"github.com/ruandg/microservices/order/internal/adapters/payment" // Novo adapter importado
	"github.com/ruandg/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// Inicializa o adaptador que liga para o microserviço de pagamentos
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	// Passa tanto o banco quanto o pagamento para a inteligência da aplicação
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	grpcAdapter.Run()
}