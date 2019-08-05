package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestXML_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.XML{}, openapi.ErrRequired{Target: "xml.namespace"}},
		{"invalidURLNamespace", openapi.XML{Namespace: "foobar"}, openapi.ErrFormatInvalid{Target: "xml.namespace", Format: "URL"}},
		{"withNamespace", openapi.XML{Namespace: exampleCom}, nil},
	}
	testValidater(t, candidates)
}
