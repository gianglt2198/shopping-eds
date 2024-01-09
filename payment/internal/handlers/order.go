package handlers

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/order/orderspb"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	eventMsghandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eMsg am.EventMessage) error {
		return orderHandlers.HandleEvent(ctx, eMsg)
	})

	return stream.Subscribe(orderspb.OrderAggregateChannel, eventMsghandler, am.MessageFilter{
		orderspb.OrderCheckedOutEvent,
		orderspb.OrderCanceledEvent,
	})
}
