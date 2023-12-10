package registry

type Serde interface {
	Register(Registrable, ...BuildOption) error
	RegisterKey(string, interface{}, ...BuildOption) error
	RegisterFactory(string, func() interface{}, ...BuildOption) error
}
