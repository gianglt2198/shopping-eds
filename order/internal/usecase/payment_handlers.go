package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/usecase/commands"
	"shopping/payment/paymentspb"
)

type PaymentHandlers[T ddd.Event] struct {
	orderUsecase ServiceUsecase
}

var _ ddd.EventHandler[ddd.Event] = (*PaymentHandlers[ddd.Event])(nil)

func NewPaymentHandlers(orderUsecase ServiceUsecase) PaymentHandlers[ddd.Event] {
	return PaymentHandlers[ddd.Event]{
		orderUsecase: orderUsecase,
	}
}

func (h PaymentHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case paymentspb.InvoiceCreatedEvent:
		return h.onInvoiceCreated(ctx, event)
	case paymentspb.InvoicePaidEvent:
		return h.onInvoicePaid(ctx, event)
	}
	return nil
}

func (h PaymentHandlers[T]) onInvoiceCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*paymentspb.InvoiceCreated)
	return h.orderUsecase.ReadyOrder(ctx, commands.ReadyOrder{
		ID:        payload.GetOrderId(),
		PaymentID: payload.GetId(),
	})
}

func (h PaymentHandlers[T]) onInvoicePaid(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*paymentspb.InvoicePaid)
	return h.orderUsecase.CompleteOrder(ctx, commands.CompleteOrder{
		ID: payload.GetOrderId(),
	})
}
