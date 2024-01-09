package usecase

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
	"shopping/order/orderspb"
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
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case domain.OrderAddedItemEvent:
		return h.onOrderAddedItem(ctx, event)
	case domain.OrderCheckedOutEvent:
		return h.onOrderCheckedOut(ctx, event)
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case domain.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case domain.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	}

	return nil
}

func (h IntegrationEventHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderCreated)

	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderCreatedEvent, &orderspb.OrderCreated{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
		}))
}

func (h IntegrationEventHandlers[T]) onOrderAddedItem(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderAddedItem)
	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderAddedItemEvent, &orderspb.OrderAddedItem{
			Id: event.AggregateID(),
			Item: &orderspb.OrderAddedItem_Item{
				ProductId: payload.Item.ProductID,
				Price:     payload.Item.Price,
				Quantity:  payload.Item.Quantity,
			},
		}))
}

func (h IntegrationEventHandlers[T]) onOrderCheckedOut(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderCheckedout)
	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderCheckedOutEvent, &orderspb.OrderCheckedOut{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
			Total:      payload.Total,
		}))
}

func (h IntegrationEventHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderReadied)
	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderReadiedEvent, &orderspb.OrderReadied{
			Id:        event.AggregateID(),
			PaymentId: payload.PaymenID,
		}))
}

func (h IntegrationEventHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderCancelled)
	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderCanceledEvent, &orderspb.OrderCanceled{
			Id:        event.AggregateID(),
			PaymentId: payload.PaymentID,
		}))
}

func (h IntegrationEventHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	return h.publisher.Publish(ctx, orderspb.OrderAggregateChannel,
		ddd.NewEvent(orderspb.OrderCompletedEvent, &orderspb.OrderCompleted{
			Id: event.AggregateID(),
		}))
}
