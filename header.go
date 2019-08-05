package oas3

// Header Object
type Header struct {
	Description     string `yaml:",omitempty"`
	Required        bool   `yaml:",omitempty"`
	Deprecated      string `yaml:",omitempty"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue,omitempty"`

	Style         string              `yaml:",omitempty"`
	Explode       bool                `yaml:",omitempty"`
	AllowReserved bool                `yaml:"allowReserved,omitempty"`
	Schema        *Schema             `yaml:",omitempty"`
	Example       interface{}         `yaml:",omitempty"`
	Examples      map[string]*Example `yaml:",omitempty"`

	Content map[string]*MediaType `yaml:",omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// Validate the values of Header object.
func (header Header) Validate() error {
	validaters := []validater{}
	if header.Schema != nil {
		validaters = append(validaters, header.Schema)
	}
	if v, ok := header.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	if len(header.Content) > 1 {
		return ErrTooManyHeaderContent
	}
	for _, mediaType := range header.Content {
		validaters = append(validaters, mediaType)
	}
	return validateAll(validaters)
}
