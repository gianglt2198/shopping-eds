package domain

import (
	"github.com/stackus/errors"
)

type Invoice struct {
	ID         string
	OrderID    string
	CustomerID string
	Amount     float64
	Status     PaymentStatus
}

func CreateInvoice(id, orderID, customerID string, amount float64) (*Invoice, error) {
	if id == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the invoice id cannot be blank")
	}

	if orderID == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the order id cannot be blank")
	}

	if customerID == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	}

	if amount <= 0 {
		return nil, errors.Wrap(errors.ErrBadRequest, "the amount must be greater than zero")
	}

	return &Invoice{
		ID:         id,
		OrderID:    orderID,
		CustomerID: customerID,
		Amount:     amount,
		Status:     PaymentPending,
	}, nil
}

func (o *Invoice) Cancel() error {
	o.Status = PaymentCancelled

	return nil
}

func (o *Invoice) Pay() error {
	o.Status = PaymentPaid

	return nil
}
