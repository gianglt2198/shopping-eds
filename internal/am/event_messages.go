package am

import (
	"context"
	"shopping/internal/ddd"
	"shopping/internal/registry"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	EventMessage interface {
		Message
		ddd.Event
	}

	EventPublisher  = MessagePublisher[ddd.Event]
	EventSubscriber = MessageSubscriber[EventMessage]
	EventStream     = MessageStream[ddd.Event, EventMessage]

	eventStream struct {
		reg    registry.Registry
		stream MessageStream[RawMessage, RawMessage]
	}

	eventMessage struct {
		id         string
		name       string
		payload    ddd.EventPayload
		metadata   ddd.Metadata
		occurredAt time.Time
		msg        RawMessage
	}
)

var (
	_ EventMessage = (*eventMessage)(nil)
	_ EventStream  = (*eventStream)(nil)
)

func NewEventStream(reg registry.Registry, stream MessageStream[RawMessage, RawMessage]) EventStream {
	return &eventStream{
		reg:    reg,
		stream: stream,
	}
}

func (e eventStream) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	metadata, err := structpb.NewStruct(event.Metadata())
	if err != nil {
		return err
	}

	payload, err := e.reg.Serialize(event.EventName(), event.Payload())
	if err != nil {
		return err
	}

	data, err := proto.Marshal(&EventMessageData{
		Payload:    payload,
		Metadata:   metadata,
		OccurredAt: timestamppb.New(event.OccurredAt()),
	})
	if err != nil {
		return err
	}

	return e.stream.Publish(ctx, topicName, rawMessage{
		id:   event.ID(),
		name: event.EventName(),
		data: data,
	})
}

func (e eventStream) Subscribe(topicName string, handler MessageHandler[EventMessage], ops ...SubscriberOption) error {
	cfg := NewSubscriberConfig(ops)

	var filters map[string]interface{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]interface{})
		for _, key := range cfg.MessageFilters() {
			filters[key] = struct{}{}
		}
	}

	fn := MessageHandlerFunc[RawMessage](func(ctx context.Context, msg RawMessage) error {
		var eventData EventMessageData

		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		err := proto.Unmarshal(msg.Data(), &eventData)
		if err != nil {
			return err
		}

		eventName := msg.MessageName()

		payload, err := e.reg.Deserialize(eventName, eventData.GetPayload())
		if err != nil {
			return err
		}

		eventMsg := eventMessage{
			id:         msg.ID(),
			name:       eventName,
			payload:    payload,
			metadata:   eventData.GetMetadata().AsMap(),
			occurredAt: eventData.GetOccurredAt().AsTime(),
			msg:        msg,
		}

		return handler.HandleMessage(ctx, eventMsg)
	})

	return e.stream.Subscribe(topicName, fn, ops...)
}

func (e eventMessage) ID() string                { return e.id }
func (e eventMessage) EventName() string         { return e.name }
func (e eventMessage) Payload() ddd.EventPayload { return e.payload }
func (e eventMessage) Metadata() ddd.Metadata    { return e.metadata }
func (e eventMessage) OccurredAt() time.Time     { return e.occurredAt }
func (e eventMessage) MessageName() string       { return e.msg.MessageName() }
func (e eventMessage) Ack() error                { return e.msg.Ack() }
func (e eventMessage) NAck() error               { return e.msg.NAck() }
func (e eventMessage) Extend() error             { return e.msg.Extend() }
func (e eventMessage) Kill() error               { return e.msg.Kill() }
