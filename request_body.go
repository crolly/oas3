package oas3

// RequestBody Object
type RequestBody struct {
	Description string                `yaml:",omitempty"`
	Content     map[string]*MediaType `yaml:",omitempty"`
	Required    bool                  `yaml:",omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// Validate the values of RequestBody object.
func (requestBody RequestBody) Validate() error {
	if requestBody.Ref != "" {
		return nil // validated in doc.Components
	}
	if requestBody.Content == nil || len(requestBody.Content) == 0 {
		return ErrRequired{Target: "requestBody.content"}
	}
	for _, mediaType := range requestBody.Content {
		if err := mediaType.Validate(); err != nil {
			return err
		}
	}
	return nil
}
