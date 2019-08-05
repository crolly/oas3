package oas3

// ServerVariable Object
type ServerVariable struct {
	Enum        []string `yaml:",omitempty"`
	Default     string   `yaml:",omitempty"`
	Description string   `yaml:",omitempty"`
}

// Validate the values of Server Variable object.
func (sv ServerVariable) Validate() error {
	if sv.Default == "" {
		return ErrRequired{Target: "serverVariable.default"}
	}
	return nil
}
