package grpc

import (
	"context"
	"fmt"

	"github.com/ruandg/microservices-proto/golang/shipping"
	"github.com/ruandg/microservices/shipping/internal/ports"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api ports.APIPort
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (a *Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipment...")

	// Soma a quantidade de todos os itens pra calcular o prazo de entrega
	var totalUnits int32
	for _, item := range request.Items {
		totalUnits += item.Quantity
	}

	_, deliveryDays, err := a.api.Create(ctx, request.OrderId, totalUnits)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to create shipment. %v", err)).Err()
	}

	return &shipping.CreateShippingResponse{DeliveryDays: deliveryDays}, nil
}