package handlers

import (
	"context"
	"shopping/customer/customerspb"
	"shopping/internal/am"
	"shopping/internal/ddd"
)

func RegisterCustomerHandlers(customerHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	eventMsghandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eMsg am.EventMessage) error {
		return customerHandlers.HandleEvent(ctx, eMsg)
	})

	return stream.Subscribe(customerspb.CustomerAggregateChannel, eventMsghandler, am.MessageFilter{
		customerspb.CustomerRegisteredEvent,
	})
}
