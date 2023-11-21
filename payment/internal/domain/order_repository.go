package domain

import "context"

type OrderRepository interface {
	CompleteOrder(ctx context.Context, orderID string) error
}
