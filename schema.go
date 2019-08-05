package oas3

// Schema Object
type Schema struct {
	Title            string   `yaml:",omitempty"`
	MultipleOf       int      `yaml:"multipleOf,omitempty"`
	Maximum          int      `yaml:",omitempty"`
	ExclusiveMaximum bool     `yaml:"exclusiveMaximum,omitempty"`
	Minimum          int      `yaml:",omitempty"`
	ExclusiveMinimum bool     `yaml:"exclusiveMinimum,omitempty"`
	MaxLength        int      `yaml:"maxLength,omitempty"`
	MinLength        int      `yaml:"minLength,omitempty"`
	Pattern          string   `yaml:",omitempty"`
	MaxItems         int      `yaml:"maxItems,omitempty"`
	MinItems         int      `yaml:"minItems,omitempty"`
	MaxProperties    int      `yaml:"maxProperties,omitempty"`
	MinProperties    int      `yaml:"minProperties,omitempty"`
	Required         []string `yaml:",omitempty"`
	Enum             []string `yaml:",omitempty"`

	Type                 string             `yaml:",omitempty"`
	AllOf                []*Schema          `yaml:"allOf,omitempty"`
	OneOf                []*Schema          `yaml:"oneOf,omitempty"`
	AnyOf                []*Schema          `yaml:"anyOf,omitempty"`
	Not                  *Schema            `yaml:",omitempty"`
	Items                *Schema            `yaml:",omitempty"`
	Properties           map[string]*Schema `yaml:",omitempty"`
	AdditionalProperties *Schema            `yaml:"additionalProperties,omitempty"`
	Description          string             `yaml:",omitempty"`
	Format               string             `yaml:",omitempty"`
	Default              string             `yaml:",omitempty"`

	Nullable      bool                   `yaml:",omitempty"`
	Discriminator *Discriminator         `yaml:",omitempty"`
	ReadOnly      bool                   `yaml:"readOnly,omitempty"`
	WriteOnly     bool                   `yaml:"writeOnly,omitempty"`
	XML           *XML                   `yaml:",omitempty"`
	ExternalDocs  *ExternalDocumentation `yaml:"externalDocs,omitempty"`
	Example       interface{}            `yaml:",omitempty"`
	Deprecated    bool                   `yaml:",omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// Validate the values of Schema object.
func (schema Schema) Validate() error {
	validaters := []validater{}
	for _, s := range schema.AllOf {
		validaters = append(validaters, s)
	}
	for _, s := range schema.OneOf {
		validaters = append(validaters, s)
	}
	for _, s := range schema.AnyOf {
		validaters = append(validaters, s)
	}
	if schema.Not != nil {
		validaters = append(validaters, schema.Not)
	}
	if schema.Items != nil {
		validaters = append(validaters, schema.Items)
	}
	if schema.Discriminator != nil {
		validaters = append(validaters, schema.Discriminator)
	}
	if schema.XML != nil {
		validaters = append(validaters, schema.XML)
	}
	if schema.ExternalDocs != nil {
		validaters = append(validaters, schema.ExternalDocs)
	}
	for _, property := range schema.Properties {
		validaters = append(validaters, property)
	}
	if e, ok := schema.Example.(validater); ok {
		validaters = append(validaters, e)
	}
	return validateAll(validaters)
}
