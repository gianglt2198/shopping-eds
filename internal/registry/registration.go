package registry

import (
	"fmt"
	"reflect"
)

func Register(r Registry, v Registrable, s Serializer, d Deserializer, ops []BuildOption) error {
	var key string

	t := reflect.TypeOf(v)

	// 	if type of t is (*T)(nil), value will be accepted
	//  apart from that, everything is acceped

	switch {
	case t.Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil():
		key = reflect.New(t).Interface().(Registrable).Key()
	default:
		key = v.Key()
	}

	return RegisterKey(r, key, v, s, d, ops)
}

func RegisterKey(r Registry, key string, v interface{}, s Serializer, d Deserializer, ops []BuildOption) error {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return r.register(key, func() interface{} {
		return reflect.New(t).Interface()
	}, s, d, ops)
}

func RegisterFactory(r Registry, key string, fn func() interface{}, s Serializer, d Deserializer, ops []BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("factory for item `%s` returns a nil value", key)
	}

	if t := reflect.TypeOf(fn()); t.Kind() != reflect.Ptr {
		return fmt.Errorf("factory for item `%s` does not return a pointer receiver", key)
	}

	return r.register(key, fn, s, d, ops)
}
