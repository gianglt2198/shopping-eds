package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

type PaymentHandlers[T ddd.AggregateEvent] struct {
	payments domain.PaymentRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*PaymentHandlers[ddd.AggregateEvent])(nil)

func NewPaymentHandlers(payments domain.PaymentRepository) *PaymentHandlers[ddd.AggregateEvent] {
	return &PaymentHandlers[ddd.AggregateEvent]{
		payments: payments,
	}
}

func (h PaymentHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.OrderCheckedOutEvent:
		return h.OnOrderCheckedout(ctx, event)
	case domain.OrderReadiedEvent:
		return h.OnOrderCancelled(ctx, event)
	}
	return nil
}

func (h PaymentHandlers[T]) OnOrderCheckedout(ctx context.Context, event ddd.AggregateEvent) error {
	orderCheckedout := event.Payload().(*domain.OrderCheckedout)
	_, err := h.payments.CreateInvoice(ctx, event.AggregateID(), orderCheckedout.CustomerID, orderCheckedout.Total)
	return err
}

func (h PaymentHandlers[T]) OnOrderCancelled(ctx context.Context, event ddd.AggregateEvent) error {
	orderCancelled := event.Payload().(*domain.OrderCancelled)
	if orderCancelled.PaymentID != "" {
		return h.payments.CancelInvoice(ctx, orderCancelled.PaymentID)
	}
	return nil
}
