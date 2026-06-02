package ports

import "github.com/ruandg/microservices/order/internal/application/core/domain"

type DBPort interface {
	Save(order *domain.Order) error
	Get(id string) (domain.Order, error)
}