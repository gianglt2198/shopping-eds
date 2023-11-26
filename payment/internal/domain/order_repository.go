package domain

import "context"

type OrderRepository interface {
	ReadyOrder(context.Context, string, string) error
	CompleteOrder(ctx context.Context, orderID string) error
}
