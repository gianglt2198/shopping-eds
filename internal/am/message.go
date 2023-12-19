package am

import (
	"context"
	"shopping/internal/ddd"
)

type (
	Message interface {
		ddd.IDer
		MessageName() string
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	MessageHandler[O Message] interface {
		HandleMessage(context.Context, O) error
	}

	MessageHandlerFunc[O Message] func(context.Context, O) error

	MessagePublisher[I any] interface {
		Publish(context.Context, string, I) error
	}

	MessageSubscriber[O Message] interface {
		Subscribe(string, MessageHandler[O], ...SubscriberOption) error
	}

	MessageStream[I any, O Message] interface {
		MessagePublisher[I]
		MessageSubscriber[O]
	}
)

func (f MessageHandlerFunc[O]) HandleMessage(ctx context.Context, msg O) error {
	return f(ctx, msg)
}
