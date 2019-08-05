package oas3

// MediaType Object
type MediaType struct {
	Schema   *Schema              `yaml:",omitempty"`
	Example  interface{}          `yaml:",omitempty"`
	Examples map[string]*Example  `yaml:",omitempty"`
	Encoding map[string]*Encoding `yaml:",omitempty"`
}

// Validate the values of MediaType object.
// This function DOES NOT check whether the encoding object is in schema or not.
func (mediaType MediaType) Validate() error {
	validaters := []validater{}
	if mediaType.Schema != nil {
		validaters = append(validaters, mediaType.Schema)
	}
	if v, ok := mediaType.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	for _, e := range mediaType.Encoding {
		validaters = append(validaters, e)
	}
	return validateAll(validaters)
}
