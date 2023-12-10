package serdes

import (
	"encoding/json"
	"shopping/internal/registry"
)

type JsonSerde struct {
	r registry.Registry
}

var _ registry.Serde = (*JsonSerde)(nil)

func NewJsonSerde(r registry.Registry) *JsonSerde {
	return &JsonSerde{
		r: r,
	}
}

func (JsonSerde) serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JsonSerde) deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (s *JsonSerde) Register(v registry.Registrable, ops ...registry.BuildOption) error {
	return registry.Register(s.r, v, s.serialize, s.deserialize, ops)
}

func (s *JsonSerde) RegisterKey(key string, v interface{}, ops ...registry.BuildOption) error {
	return registry.RegisterKey(s.r, key, v, s.serialize, s.deserialize, ops)
}

func (s *JsonSerde) RegisterFactory(key string, fn func() interface{}, ops ...registry.BuildOption) error {
	return registry.RegisterFactory(s.r, key, fn, s.serialize, s.deserialize, ops)
}
