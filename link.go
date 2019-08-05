package oas3

import "errors"

// Link Object
type Link struct {
	OperationRef string `yaml:"operationRef"`
	OperationID  string `yaml:"operationId"`
	Parameters   map[string]interface{}
	RequestBody  interface{} `yaml:"requestBody"`
	Description  string
	Server       *Server

	Ref string `yaml:"$ref"`
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
