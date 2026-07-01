// ports/api.go
package ports

import (
	"context"

	"github.com/ruandg/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	Create(ctx context.Context, orderID int64, totalUnits int32) (domain.Shipping, int32, error)
}