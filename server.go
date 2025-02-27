package oas3

import (
	"net/url"
	"regexp"
)

var tmplVarRegexp = regexp.MustCompile("{[^}]+}")

// Server Object
type Server struct {
	URL         string                     `yaml:",omitempty"`
	Description string                     `yaml:",omitempty"`
	Variables   map[string]*ServerVariable `yaml:",omitempty"`
}

// Validate the values of Server object.
func (server Server) Validate() error {
	if err := server.validateRequiredFields(); err != nil {
		return err
	}
	// replace template variable with placeholder to validate the replaced string
	// is valid URL or not
	serverURL := tmplVarRegexp.ReplaceAllLiteralString(server.URL, "ph")
	// use url.Parse because relative URL is allowed
	if _, err := url.Parse(serverURL); err != nil {
		return ErrFormatInvalid{Target: "server.url", Format: "URL"}
	}
	validaters := []validater{}
	for _, sv := range server.Variables {
		validaters = append(validaters, sv)
	}
	return validateAll(validaters)
}

func (server Server) validateRequiredFields() error {
	if server.URL == "" {
		return ErrRequired{Target: "server.url"}
	}
	return nil
}
