package domain

import "context"

type ProductRepository interface {
	GetProduct(ctx context.Context, productID string) (*Product, error)
}
