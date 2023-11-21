package domain

import "context"

type ProductRepository interface {
	Save(context.Context, *Product) error
	Find(context.Context, string) (*Product, error)
	Update(context.Context, *Product) error
	Delete(context.Context, string) error
}
