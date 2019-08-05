package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestServerVariable_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.ServerVariable{}, openapi.ErrRequired{Target: "serverVariable.default"}},
		{"withDefault", openapi.ServerVariable{Default: "default"}, nil},
	}
	testValidater(t, candidates)
}
