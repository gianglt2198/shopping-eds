package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
	"shopping/product/productspb"
)

type ProductHandlers[T ddd.Event] struct {
	cache domain.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*ProductHandlers[ddd.Event])(nil)

func NewProductHandlers(cache domain.ProductCacheRepository) ProductHandlers[ddd.Event] {
	return ProductHandlers[ddd.Event]{
		cache: cache,
	}
}

func (h ProductHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case productspb.ProductCreatedEvent:
		return h.onProductCreated(ctx, event)
	case productspb.ProductPriceIncreasedEvent:
		return h.onProductPriceChanged(ctx, event)
	case productspb.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)
	case productspb.ProductDeletedEvent:
		return h.onProductDeleted(ctx, event)
	}
	return nil
}

func (h ProductHandlers[T]) onProductCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*productspb.ProductCreated)
	return h.cache.Add(ctx, payload.GetId(), payload.GetName(), payload.GetPrice())
}

func (h ProductHandlers[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*productspb.ProductPriceChanged)
	return h.cache.UpdatePrice(ctx, payload.GetId(), payload.GetDelta())
}

func (h ProductHandlers[T]) onProductDeleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*productspb.ProductDeleted)
	return h.cache.Remove(ctx, payload.GetId())
}
