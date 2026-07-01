package api

import (
	"context"

	"github.com/ruandg/microservices/shipping/internal/application/core/domain"
	"github.com/ruandg/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Create(ctx context.Context, orderID int64, totalUnits int32) (domain.Shipping, int32, error) {
	shipping := domain.NewShipping(orderID)

	err := a.db.Save(ctx, &shipping)
	if err != nil {
		return domain.Shipping{}, 0, err
	}

	deliveryDays := calculateDeliveryDays(totalUnits)

	return shipping, deliveryDays, nil
}

func calculateDeliveryDays(totalUnits int32) int32 {
	return 1 + (totalUnits / 5)
}