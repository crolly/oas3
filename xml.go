package oas3

// XML Object
type XML struct {
	Name      string `yaml:",omitempty"`
	Namespace string `yaml:",omitempty"`
	Prefix    string `yaml:",omitempty"`
	Attribute bool   `yaml:",omitempty"`
	Wrapped   bool   `yaml:",omitempty"`
}

// Validate the values of XML object.
func (xml XML) Validate() error {
	return mustURL("xml.namespace", xml.Namespace)
}
