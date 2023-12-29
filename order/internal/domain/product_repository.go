package domain

import (
	"context"
	"shopping/order/internal/models"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, productID string) (*models.Product, error)
}
