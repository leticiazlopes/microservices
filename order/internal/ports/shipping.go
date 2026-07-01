package ports

import "github.com/ruandg/microservices/order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipment(order *domain.Order) error
}