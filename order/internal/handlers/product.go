package handlers

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/product/productspb"
)

func RegisterProductHandlers(productHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	eventMsghandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eMsg am.EventMessage) error {
		return productHandlers.HandleEvent(ctx, eMsg)
	})

	return stream.Subscribe(productspb.ProductAggregateChannel, eventMsghandler, am.MessageFilter{
		productspb.ProductCreatedEvent,
		productspb.ProductPriceIncreasedEvent,
		productspb.ProductPriceDecreasedEvent,
		productspb.ProductDeletedEvent,
	})
}
