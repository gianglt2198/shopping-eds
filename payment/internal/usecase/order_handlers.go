package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/orderspb"

	"github.com/google/uuid"
)

type OrderHandlers[T ddd.Event] struct {
	paymentUsecase ServiceUsecase
}

var _ ddd.EventHandler[ddd.Event] = (*OrderHandlers[ddd.Event])(nil)

func NewOrderHandlers(paymentUsecase ServiceUsecase) *OrderHandlers[ddd.Event] {
	return &OrderHandlers[ddd.Event]{
		paymentUsecase: paymentUsecase,
	}
}

func (h OrderHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderspb.OrderCheckedOutEvent:
		return h.onOrderCheckedout(ctx, event)
	case orderspb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}
	return nil
}

func (h OrderHandlers[T]) onOrderCheckedout(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderspb.OrderCheckedOut)
	id := uuid.New().String()
	return h.paymentUsecase.CreateInvoice(ctx, CreateInvoice{
		ID:         id,
		OrderID:    payload.GetId(),
		CustomerID: payload.GetCustomerId(),
		Amount:     payload.GetTotal(),
	})
}

func (h OrderHandlers[T]) onOrderCanceled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderspb.OrderCanceled)
	return h.paymentUsecase.CancelInvoice(ctx, CancelInvoice{
		ID: payload.GetPaymentId(),
	})
}
