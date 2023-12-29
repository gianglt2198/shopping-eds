package domain

import (
	"context"
	"shopping/order/internal/models"
)

type CustomerRepository interface {
	GetCustomer(ctx context.Context, customerID string) (*models.Customer, error)
}
