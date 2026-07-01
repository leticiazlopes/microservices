package shipping

import (
	"context"
	"time"

	"github.com/ruandg/microservices-proto/golang/shipping"
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) CreateShipment(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var items []*shipping.ShippingItem
	for _, item := range order.OrderItems {
		items = append(items, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	_, err := a.shipping.Create(ctx, &shipping.CreateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})
	return err
}