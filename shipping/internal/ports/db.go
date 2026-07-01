package ports

import (
	"context"
	"github.com/ruandg/microservices/shipping/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Shipping, error)
	Save(ctx context.Context, shipping *domain.Shipping) error
}