package logging

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/usecase"

	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	usecase.DomainEventHandlers
	logger zerolog.Logger
}

var _ usecase.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers usecase.DomainEventHandlers, logger zerolog.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnOrderCheckedout(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Order.OnOrderCheckedout")
	defer func() { h.logger.Info().Err(err).Msg("<-- Order.OnOrderCheckedout") }()
	return h.DomainEventHandlers.OnOrderCheckedout(ctx, event)
}

func (h DomainEventHandlers) OnOrderCancelled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Order.OnOrderCancelled")
	defer func() { h.logger.Info().Err(err).Msg("<-- Order.OnOrderCancelled") }()
	return h.DomainEventHandlers.OnOrderCancelled(ctx, event)
}
