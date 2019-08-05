package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestSchema_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.Schema{}, nil},
	}
	testValidater(t, candidates)
}
