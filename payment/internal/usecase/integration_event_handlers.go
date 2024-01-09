package usecase

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/payment/internal/domain"
	"shopping/payment/paymentspb"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.InvoiceCreatedEvent:
		return h.onInvoiceCreatedEvent(ctx, event)
	case domain.InvoicePaidEvent:
		return h.onInvoicePaidEvent(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onInvoiceCreatedEvent(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.InvoiceCreated)
	return h.publisher.Publish(ctx, paymentspb.PaymentAggregateChannel,
		ddd.NewEvent(paymentspb.InvoiceCreatedEvent, &paymentspb.InvoiceCreated{
			Id:      payload.ID,
			OrderId: payload.OrderID,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onInvoicePaidEvent(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.InvoicePaid)
	return h.publisher.Publish(ctx, paymentspb.PaymentAggregateChannel,
		ddd.NewEvent(paymentspb.InvoicePaidEvent, &paymentspb.InvoicePaid{
			Id:      payload.ID,
			OrderId: payload.OrderID,
		}),
	)
}
