package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestExternalDocumentation_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.ExternalDocumentation{}, openapi.ErrRequired{Target: "externalDocumentation.url"}},
		{"invalidURL", openapi.ExternalDocumentation{URL: "foobar"}, openapi.ErrFormatInvalid{Target: "externalDocumentation.url", Format: "URL"}},
		{"valid", openapi.ExternalDocumentation{URL: exampleCom}, nil},
	}
	testValidater(t, candidates)
}
