package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestHeader_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.Header{}, nil},
		{"2 contents", openapi.Header{
			Content: map[string]*openapi.MediaType{
				"application/json": &openapi.MediaType{},
				"image/png":        &openapi.MediaType{},
			},
		}, openapi.ErrTooManyHeaderContent},
	}
	testValidater(t, candidates)
}
