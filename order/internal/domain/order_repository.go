package domain

import "context"

type OrderRepository interface {
	Load(context.Context, string) (*Order, error)
	Save(context.Context, *Order) error
}
