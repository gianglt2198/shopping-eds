package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
)

type ManagementHandlers[T ddd.AggregateEvent] struct {
	management domain.ManagementRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*ManagementHandlers[ddd.AggregateEvent])(nil)

func NewManagementHandlers(management domain.ManagementRepository) *ManagementHandlers[ddd.AggregateEvent] {
	return &ManagementHandlers[ddd.AggregateEvent]{
		management: management,
	}
}

func (h ManagementHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.ProductCreatedEvent:
		return h.onProductCreated(ctx, event)
	case domain.ProductDeletedEvent:
		return h.onProductDeleted(ctx, event)
	case domain.ProductInscreasedPriceEvent:
		return h.onProductPriceChanged(ctx, event)
	case domain.ProductDescreasedPriceEvent:
		return h.onProductPriceChanged(ctx, event)
	}
	return nil
}

func (h ManagementHandlers[T]) onProductCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.ProductCreated)
	return h.management.CreateProduct(ctx, event.AggregateID(), payload.Name, payload.Description, payload.Price)
}

func (h ManagementHandlers[T]) onProductPriceChanged(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.management.UpdatePrice(ctx, event.AggregateID(), payload.Delta)
}

func (h ManagementHandlers[T]) onProductDeleted(ctx context.Context, event T) error {
	return h.management.DeleteProduct(ctx, event.AggregateID())
}
