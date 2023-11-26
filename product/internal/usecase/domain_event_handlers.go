package usecase

import (
	"context"
	"shopping/internal/ddd"
)

type DomainEventHandlers interface {
	OnProductCreated(context.Context, ddd.Event) error
	OnProductUpdated(context.Context, ddd.Event) error
	OnProductDeleted(context.Context, ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnProductCreated(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductUpdated(context.Context, ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductDeleted(context.Context, ddd.Event) error {
	return nil
}
