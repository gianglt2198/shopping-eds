package registry

import (
	"fmt"
	"sync"
)

type (
	Registrable interface {
		Key() string
	}

	Serializer   func(interface{}) ([]byte, error)
	Deserializer func([]byte, interface{}) error

	Registry interface {
		Build(string, ...BuildOption) (interface{}, error)
		Serialize(string, interface{}) ([]byte, error)
		Deserialize(string, []byte, ...BuildOption) (interface{}, error)
		register(string, func() interface{}, Serializer, Deserializer, []BuildOption) error
	}
)

type registered struct {
	factory      func() interface{}
	serializer   Serializer
	deserializer Deserializer
	option       []BuildOption
}

type registry struct {
	registered map[string]registered
	mu         sync.Mutex
}

func New() *registry {
	return &registry{
		registered: make(map[string]registered),
	}
}

func (r *registry) Build(key string, options ...BuildOption) (interface{}, error) {
	reg, ok := r.registered[key]
	if !ok {
		return nil, fmt.Errorf("nothing has been registered with the key `%s`", key)
	}

	v := reg.factory()
	ops := append(reg.option, options...)

	for _, op := range ops {
		err := op(v)
		if err != nil {
			return nil, err
		}
	}

	return v, nil
}

func (r *registry) Serialize(key string, v interface{}) ([]byte, error) {
	reg, ok := r.registered[key]
	if !ok {
		return nil, fmt.Errorf("nothing has been registered with the key `%s`", key)
	}

	return reg.serializer(v)
}

func (r *registry) Deserialize(key string, data []byte, options ...BuildOption) (interface{}, error) {
	v, err := r.Build(key, options...)
	if err != nil {
		return nil, err
	}

	err = r.registered[key].deserializer(data, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *registry) register(key string, fn func() interface{}, s Serializer, d Deserializer, ops []BuildOption) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.registered[key]; ok {
		return fmt.Errorf("something with the key `%s` has already been registered", key)
	}

	r.registered[key] = registered{
		factory:      fn,
		serializer:   s,
		deserializer: d,
		option:       ops,
	}

	return nil
}
