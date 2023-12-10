package domain

import "context"

type ProductRepository interface {
	Load(context.Context, string) (*Product, error)
	Save(context.Context, *Product) error
}
