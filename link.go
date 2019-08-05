package oas3

import "errors"

// Link Object
type Link struct {
	OperationRef string                 `yaml:"operationRef,omitempty"`
	OperationID  string                 `yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `yaml:",omitempty"`
	RequestBody  interface{}            `yaml:"requestBody,omitempty"`
	Description  string                 `yaml:",omitempty"`
	Server       *Server                `yaml:",omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// Validate the values of Link object.
func (link Link) Validate() error {
	if link.OperationRef != "" && link.OperationID != "" {
		return errors.New("operationRef and operationId are mutually exclusive")
	}
	validaters := []validater{}
	for _, i := range link.Parameters {
		if v, ok := i.(validater); ok {
			validaters = append(validaters, v)
		}
	}
	if v, ok := link.RequestBody.(validater); ok {
		validaters = append(validaters, v)
	}
	if link.Server != nil {
		validaters = append(validaters, link.Server)
	}
	return validateAll(validaters)
}
