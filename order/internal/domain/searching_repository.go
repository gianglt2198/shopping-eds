package domain

import (
	"context"
	"shopping/order/internal/models"
	"time"
)

type (
	Filters struct {
		CustomerID string
		After      time.Time
		Before     time.Time
		ProductIDs []string
		MinTotal   float64
		MaxTotal   float64
		Status     string
	}
	SearchOrders struct {
		Filters Filters
		Next    string
		Limit   int
	}
)
type SearchingRepository interface {
	Add(ctx context.Context, order *models.Order) error
	UpdateStatus(ctx context.Context, orderID, status string) error
	UpdateItem(ctx context.Context, orderID, productID string, quantity int, price float64) error
	Search(ctx context.Context, search SearchOrders) ([]*models.Order, error)
	Get(ctx context.Context, orderID string) (*models.Order, error)
}
