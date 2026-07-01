package domain

import (
	"time"
)

type Shipping struct {
	ID        int64  `json:"id"`
	OrderID   int64  `json:"order_id"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

func NewShipping(orderId int64) Shipping {
	return Shipping{
		CreatedAt: time.Now().Unix(),
		Status:    "Pending",
		OrderID:   orderId,
	}
}