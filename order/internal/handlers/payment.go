package handlers

import (
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

func RegisterPaymentHandlers(paymentHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.OrderCheckedOutEvent, paymentHandlers)
	domainSubscriber.Subscribe(domain.OrderCanceledEvent, paymentHandlers)
}
