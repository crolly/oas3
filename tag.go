package oas3

// Tag Object
type Tag struct {
	Name         string                 `yaml:",omitempty"`
	Description  string                 `yaml:",omitempty"`
	ExternalDocs *ExternalDocumentation `yaml:"externalDocs,omitempty"`
}

// Validate the values of Tag object.
func (tag Tag) Validate() error {
	if tag.Name == "" {
		return ErrRequired{Target: "tag.name"}
	}
	if tag.ExternalDocs != nil {
		return tag.ExternalDocs.Validate()
	}
	return nil
}
