package oas3_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	openapi "github.com/crolly/oas3"
)

const (
	exampleCom  = "https://example.com"
	exampleMail = "foo@example.com"
)

type candidate struct {
	label string
	in    openapi.Validater
	err   error
}

func testValidater(t *testing.T, candidates []candidate) {
	t.Helper()
	for i, c := range candidates {
		t.Run(strconv.Itoa(i)+"/"+c.label, func(t *testing.T) {
			if err := c.in.Validate(); !reflect.DeepEqual(err, c.err) {
				t.Errorf("error should be %s, but %s", c.err, err)
			}
		})
	}
}

type mockValidater struct {
	err error
}

func (v mockValidater) Validate() error {
	return v.err
}

func TestValidateAll(t *testing.T) {
	valid := mockValidater{}
	invalid := mockValidater{errors.New("err")}

	candidates := []struct {
		label  string
		in     []openapi.Validater
		hasErr bool
	}{
		{"nil", nil, false},
		{"empty", []openapi.Validater{}, false},
		{"all valid", []openapi.Validater{valid, valid, valid}, false},
		{"have invalid", []openapi.Validater{valid, invalid, valid}, true},
		{"have nil", []openapi.Validater{valid, nil, valid, valid}, false},
	}
	for i, c := range candidates {
		t.Run(strconv.Itoa(i)+"/"+c.label, func(t *testing.T) {
			if err := openapi.ValidateAll(c.in); (err != nil) != c.hasErr {
				t.Errorf("error: %s", err)
			}
		})
	}
}

func TestMustURL(t *testing.T) {
	candidates := []struct {
		label  string
		in     string
		hasErr bool
	}{
		{"empty", "", true},
		{"valid HTTP url", "http://example.com", false},
		{"allowed relative path", "foo/bar/baz", true},
		{"absolute path", "/foo/bar/baz", false},
		{"plain string", "foobarbaz", true},
	}
	for i, c := range candidates {
		t.Run(strconv.Itoa(i)+"/"+c.label, func(t *testing.T) {
			if err := openapi.MustURL(c.label, c.in); (err != nil) != c.hasErr {
				if c.hasErr {
					t.Error("error should occured, but not")
					return
				}
				t.Error("error should not occurred, but occurred")
				return
			}
		})
	}
}
