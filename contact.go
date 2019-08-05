package oas3

// Contact Object
type Contact struct {
	Name  string
	URL   string
	Email string
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
