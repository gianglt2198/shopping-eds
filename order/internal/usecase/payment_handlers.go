package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

type PaymentHandlers struct {
	payments domain.PaymentRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*PaymentHandlers)(nil)

func NewPaymentHandlers(payments domain.PaymentRepository) *PaymentHandlers {
	return &PaymentHandlers{
		payments: payments,
	}
}

func (h PaymentHandlers) OnOrderCheckedout(ctx context.Context, event ddd.Event) error {
	orderCheckedout := event.(*domain.OrderCheckedout)
	_, err := h.payments.CreateInvoice(ctx, orderCheckedout.Order.ID, orderCheckedout.Order.CustomerID, orderCheckedout.Order.GetTotal())
	return err
}

func (h PaymentHandlers) OnOrderCancelled(ctx context.Context, event ddd.Event) error {
	orderCancelled := event.(*domain.OrderCancelled)
	if orderCancelled.Order.PaymentID != "" {
		return h.payments.CancelInvoice(ctx, orderCancelled.Order.PaymentID)
	}
	return nil
}
