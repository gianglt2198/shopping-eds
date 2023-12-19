package handlers

import (
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

func RegisterPaymentHandlers(paymentHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(paymentHandlers,
		domain.OrderCheckedOutEvent,
		domain.OrderCanceledEvent,
	)
}
