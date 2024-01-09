package handlers

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/ddd"
	"shopping/payment/paymentspb"
)

func RegisterPaymentHandlers(paymentHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	eventMsghandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eMsg am.EventMessage) error {
		return paymentHandlers.HandleEvent(ctx, eMsg)
	})

	return stream.Subscribe(paymentspb.PaymentAggregateChannel, eventMsghandler, am.MessageFilter{
		paymentspb.InvoiceCreatedEvent,
		paymentspb.InvoicePaidEvent,
	})
}
