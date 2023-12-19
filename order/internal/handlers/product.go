package handlers

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/product/pb"
)

func RegisterProductHandlers(productHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	eventMsghandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eMsg am.EventMessage) error {
		return productHandlers.HandleEvent(ctx, eMsg)
	})

	return stream.Subscribe(pb.ProductAggregateChannel, eventMsghandler, am.MessageFilter{
		pb.ProductCreatedEvent,
		pb.ProductPriceIncreasedEvent,
		pb.ProductPriceDecreasedEvent,
		pb.ProductDeletedEvent,
	})
}
