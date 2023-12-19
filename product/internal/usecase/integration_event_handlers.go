package usecase

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
	"shopping/product/pb"
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
	case domain.ProductCreatedEvent:
		return h.onProductCreated(ctx, event)
	case domain.ProductInscreasedPriceEvent:
		return h.onProductIncreasedPrice(ctx, event)
	case domain.ProductDescreasedPriceEvent:
		return h.onProducDecreasedPrice(ctx, event)
	case domain.ProductDeletedEvent:
		return h.onProductDeleted(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onProductCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductCreated)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductCreatedEvent, &pb.ProductCreated{
			Id:          event.ID(),
			Name:        payload.Name,
			Description: payload.Description,
			Price:       payload.Price,
		}))
}

func (h IntegrationEventHandlers[T]) onProductIncreasedPrice(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductPriceIncreasedEvent, &pb.ProductPriceChanged{
			Id:    event.ID(),
			Delta: payload.Delta,
		}))
}

func (h IntegrationEventHandlers[T]) onProducDecreasedPrice(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductPriceDecreasedEvent, &pb.ProductPriceChanged{
			Id:    event.ID(),
			Delta: payload.Delta,
		}))
}

func (h IntegrationEventHandlers[T]) onProductDeleted(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductDeletedEvent, &pb.ProductCreated{
			Id: event.ID(),
		}))
}
