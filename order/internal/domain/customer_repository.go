package domain

import (
	"context"
)

type CustomerRepository interface {
	GetCustomer(ctx context.Context, customerID string) error
}
