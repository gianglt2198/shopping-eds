package domain

import "context"

type ProductCacheRepository interface {
	Add(ctx context.Context, productID, name string, price float64) error
	UpdatePrice(ctx context.Context, productID string, delta float64) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}
