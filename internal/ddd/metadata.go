package ddd

type Metadata map[string]interface{}

func (m Metadata) Set(key string, value interface{}) {
	m[key] = value
}

func (m Metadata) Get(key string) interface{} {
	return m[key]
}

func (m Metadata) Delete(key string) {
	delete(m, key)
}

func (m Metadata) configureEvent(e *event) {
	for k, v := range m {
		e.metadata[k] = v
	}
}
