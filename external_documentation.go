package oas3

// ExternalDocumentation Object
type ExternalDocumentation struct {
	Description string `yaml:",omitempty"`
	URL         string `yaml:",omitempty"`
}

// Validate the values of ExternalDocumentaion object.
func (externalDocumentation ExternalDocumentation) Validate() error {
	return mustURL("externalDocumentation.url", externalDocumentation.URL)
}
