package oas3

// Response Object
type Response struct {
	Description string                `yaml:",omitempty"`
	Headers     map[string]*Header    `yaml:",omitempty"`
	Content     map[string]*MediaType `yaml:",omitempty"`
	Links       map[string]*Link      `yaml:",omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// Validate the value of Response object.
func (response Response) Validate() error {
	if response.Ref != "" {
		return nil // validated in doc.Components
	}
	if response.Description == "" {
		return ErrRequired{Target: "response.description"}
	}
	validaters := []validater{}
	for _, header := range response.Headers {
		validaters = append(validaters, header)
	}
	for _, mediaType := range response.Content {
		validaters = append(validaters, mediaType)
	}
	for _, link := range response.Links {
		validaters = append(validaters, link)
	}
	return validateAll(validaters)
}
