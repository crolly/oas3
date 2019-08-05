package oas3

// Example Object
type Example struct {
	Summary       string      `yaml:",omitempty"`
	Description   string      `yaml:",omitempty"`
	Value         interface{} `yaml:",omitempty"`
	ExternalValue interface{} `yaml:"externalValue,omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}
