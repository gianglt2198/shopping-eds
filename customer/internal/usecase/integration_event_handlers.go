package usecase

import (
	"context"
	"shopping/customer/customerspb"
	"shopping/customer/internal/domain"
	"shopping/internal/am"
	"shopping/internal/ddd"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	}

	return nil
}

func (h IntegrationEventHandlers[T]) onCustomerRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.CustomerRegistered)
	return h.publisher.Publish(ctx, customerspb.CustomerAggregateChannel,
		ddd.NewEvent(customerspb.CustomerRegisteredEvent, &customerspb.CustomerRegistered{
			Id:        payload.Customer.ID(),
			Name:      payload.Customer.Name,
			SmsNumber: payload.Customer.SmsNumber,
			Email:     payload.Customer.Email,
		}))
}
