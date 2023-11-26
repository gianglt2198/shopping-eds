package domain

import "context"

type PaymentRepository interface {
	Save(context.Context, *Invoice) error
	Find(context.Context, string) (*Invoice, error)
	Update(context.Context, *Invoice) error
	Delete(context.Context, *Invoice) error
}
