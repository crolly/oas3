package oas3

// Example Object
type Example struct {
	Summary       string
	Description   string
	Value         interface{}
	ExternalValue interface{} `yaml:"externalValue"`

	Ref string `yaml:"$ref"`
}
