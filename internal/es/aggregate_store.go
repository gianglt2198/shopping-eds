package es

import (
	"context"
	"shopping/internal/ddd"
)

type EventSourcedAggregate interface {
	ddd.IDer
	AggregateName() string
	ddd.Eventer
	Versioner
	EventApplier
	EventCommitter
}

type AggregateStoreMiddleware func(AggregateStore) AggregateStore

type AggregateStore interface {
	Load(context.Context, EventSourcedAggregate) error
	Save(context.Context, EventSourcedAggregate) error
}

func AggreagteStoreWithMiddleware(store AggregateStore, middlewares ...AggregateStoreMiddleware) AggregateStore {
	s := store
	// middleware are applied in reverse; this makes the first middleware
	// in the slice the outermost i.e. first to enter, last to exit
	// given: store, A, B, C
	// result: A(B(C(store)))
	for i := len(middlewares) - 1; i >= 0; i-- {
		s = middlewares[i](s)
	}

	return s
}
