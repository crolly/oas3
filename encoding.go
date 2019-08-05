package oas3

// Encoding Object
type Encoding struct {
	ContentType   string             `yaml:"contentType,omitempty"`
	Headers       map[string]*Header `yaml:",omitempty"`
	Style         string             `yaml:",omitempty"`
	Explode       bool               `yaml:",omitempty"`
	AllowReserved bool               `yaml:"allowReserved,omitempty"`
}

// Validate the values of Encoding object.
func (encoding Encoding) Validate() error {
	for _, header := range encoding.Headers {
		if err := header.Validate(); err != nil {
			return err
		}
	}
	return nil
}
