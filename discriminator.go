package oas3

// Discriminator Object
type Discriminator struct {
	PropertyName string            `yaml:"propertyName,omitempty"`
	Mapping      map[string]string `yaml:",omitempty"`
}

// Validate the values of Descriminator object.
func (discriminator Discriminator) Validate() error {
	if discriminator.PropertyName == "" {
		return ErrRequired{Target: "discriminator.propertyName"}
	}
	return nil
}
