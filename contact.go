package oas3

// Contact Object
type Contact struct {
	Name  string `yaml:",omitempty"`
	URL   string `yaml:",omitempty"`
	Email string `yaml:",omitempty"`
}

// Validate the values of Contact object.
func (contact Contact) Validate() error {
	if err := mustURL("contact.url", contact.URL); err != nil {
		return err
	}
	if contact.Email != "" {
		if !emailRegexp.MatchString(contact.Email) {
			return ErrFormatInvalid{Target: "contact.email", Format: "email"}
		}
	}
	return nil
}
