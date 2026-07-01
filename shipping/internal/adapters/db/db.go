package db

import (
	"context"
	"fmt"
	"github.com/ruandg/microservices/shipping/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Entidade que representa a tabela de Frete no banco de dados
type Shipping struct {
	gorm.Model
	OrderID int64
	Status  string
}

type Adapter struct {
	db *gorm.DB
}

// O método Get agora aceita int64 para bater com o que definimos na porta (ports.DBPort)
func (a Adapter) Get(ctx context.Context, id int64) (domain.Shipping, error) {
	var shippingEntity Shipping
	res := a.db.WithContext(ctx).First(&shippingEntity, id)
	
	shipping := domain.Shipping{
		ID:        int64(shippingEntity.ID),
		OrderID:   shippingEntity.OrderID,
		Status:    shippingEntity.Status,
		CreatedAt: shippingEntity.CreatedAt.Unix(),
	}
	return shipping, res.Error
}

func (a Adapter) Save(ctx context.Context, shipping *domain.Shipping) error {
	shippingModel := Shipping{
		OrderID: shipping.OrderID,
		Status:  shipping.Status,
	}
	res := a.db.WithContext(ctx).Create(&shippingModel)
	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// Mudamos o nome do DB no rastreamento para shipping
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("shipping"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	// Migração automática cria a tabela correta de frete
	err := db.AutoMigrate(&Shipping{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}