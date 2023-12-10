package serdes

import (
	"fmt"
	"reflect"
	"shopping/internal/registry"

	"google.golang.org/protobuf/proto"
)

type ProtoSerde struct {
	r registry.Registry
}

var _ registry.Serde = (*ProtoSerde)(nil)
var protoT = reflect.TypeOf((*proto.Message)(nil)).Elem()

func NewProtoSerde(r registry.Registry) *ProtoSerde {
	return &ProtoSerde{
		r: r,
	}
}

func (ProtoSerde) serialize(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (ProtoSerde) deserialize(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (s *ProtoSerde) Register(v registry.Registrable, ops ...registry.BuildOption) error {
	if reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.Register(s.r, v, s.serialize, s.deserialize, ops)
}

func (s *ProtoSerde) RegisterKey(key string, v interface{}, ops ...registry.BuildOption) error {
	if reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.RegisterKey(s.r, key, v, s.serialize, s.deserialize, ops)
}

func (s *ProtoSerde) RegisterFactory(key string, fn func() interface{}, ops ...registry.BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("%s factory returns a nil value", key)
	} else if _, ok := v.(proto.Message); !ok {
		return fmt.Errorf("%s does not implement proto.Message", key)
	}
	return registry.RegisterFactory(s.r, key, fn, s.serialize, s.deserialize, ops)
}
