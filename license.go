package oas3

import (
	"net/url"
)

// License Object
type License struct {
	Name string `yaml:",omitempty"`
	URL  string `yaml:",omitempty"`
}

// Validate the values of License object.
func (license License) Validate() error {
	if license.Name == "" {
		return ErrRequired{Target: "license.name"}
	}
	if license.URL != "" {
		_, err := url.ParseRequestURI(license.URL)
		if err != nil {
			return ErrFormatInvalid{Target: "license.url", Format: "URL"}
		}
		return nil
	}
	return nil
}
