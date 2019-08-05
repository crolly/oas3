package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestDiscriminator_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.Discriminator{}, openapi.ErrRequired{Target: "discriminator.propertyName"}},
		{"withPropertyName", openapi.Discriminator{PropertyName: "foobar"}, nil},
	}
	testValidater(t, candidates)
}
