package domain

import (
	"context"
)

type PaymentRepository interface {
	GetInvoice(context.Context, string) error
	CreateInvoice(ctx context.Context, orderID, customerID string, amount float64) (string, error)
	CancelInvoice(context.Context, string) error
}
