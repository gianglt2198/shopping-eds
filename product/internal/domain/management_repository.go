package domain

import "context"

type ManagementProduct struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

type ManagementRepository interface {
	CreateProduct(ctx context.Context, productID, name, description string, price float64) error
	UpdatePrice(ctx context.Context, productID string, delta float64) error
	DeleteProduct(ctx context.Context, productID string) error
	Find(ctx context.Context, productID string) (*ManagementProduct, error)
}
