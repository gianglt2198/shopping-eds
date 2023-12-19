package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/pb"

	"github.com/rs/zerolog"
)

type ProductHandlers[T ddd.Event] struct {
	logger zerolog.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*ProductHandlers[ddd.Event])(nil)

func NewProductHandlers(logger zerolog.Logger) ProductHandlers[ddd.Event] {
	return ProductHandlers[ddd.Event]{
		logger: logger,
	}
}

func (h ProductHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case pb.ProductCreatedEvent:
		return h.onProductCreated(ctx, event)
	case pb.ProductPriceIncreasedEvent:
		return h.onProductPriceIncreased(ctx, event)
	case pb.ProductPriceDecreasedEvent:
		return h.onProductPriceDecreased(ctx, event)
	case pb.ProductDeletedEvent:
		return h.onProductDeleted(ctx, event)
	}
	return nil
}

func (h ProductHandlers[T]) onProductCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*pb.ProductCreated)
	h.logger.Debug().Msgf("ID: %v, Name: %v, Description: %v, Price: %v", payload.GetId(), payload.GetName(), payload.GetDescription(), payload.GetPrice())
	return nil
}

func (h ProductHandlers[T]) onProductPriceIncreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*pb.ProductPriceChanged)
	h.logger.Debug().Msgf("ID: %v, Delta: %v", payload.GetId(), payload.GetDelta())
	return nil
}

func (h ProductHandlers[T]) onProductPriceDecreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*pb.ProductPriceChanged)
	h.logger.Debug().Msgf("ID: %v, Delta: %v", payload.GetId(), payload.GetDelta())
	return nil
}

func (h ProductHandlers[T]) onProductDeleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*pb.ProductDeleted)
	h.logger.Debug().Msgf("ID: %v", payload.GetId())
	return nil
}
