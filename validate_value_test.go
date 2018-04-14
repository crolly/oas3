package openapi

import (
	"testing"
)

func TestHasDuplicatedParameter(t *testing.T) {
	t.Run("no duplicated param", testHasDuplicatedParameterFalse)
	t.Run("there's duplicated param", testHasDuplicatedParameterTrue)
}

func testHasDuplicatedParameterFalse(t *testing.T) {
	params := []*Parameter{
		&Parameter{Name: "foo", In: "header"},
		&Parameter{Name: "foo", In: "path", Required: true},
		&Parameter{Name: "bar", In: "path", Required: true},
	}
	if hasDuplicatedParameter(params) {
		t.Error("should return false")
	}
}

func testHasDuplicatedParameterTrue(t *testing.T) {
	params := []*Parameter{
		&Parameter{Name: "foo", In: "header"},
		&Parameter{Name: "foo", In: "header"},
	}
	if !hasDuplicatedParameter(params) {
		t.Error("should return true")
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
	for _, c := range candidates {
		if err := mustURL(c.label, c.in); (err != nil) != c.hasErr {
			t.Logf("error occured at %s", c.label)
			if c.hasErr {
				t.Error("error should occured, but not")
				return
			}
			t.Error("error should not occurred, but occurred")
			return
		}
	}
}

const exampleCom = "https://example.com"

var mockScopes = map[string]string{"foo": "bar"}

func TestOAuthFlowValidate(t *testing.T) {
	empty := OAuthFlow{}
	aURL := OAuthFlow{AuthorizationURL: exampleCom}
	tURL := OAuthFlow{TokenURL: exampleCom}
	rURL := OAuthFlow{RefreshURL: exampleCom}
	scopes := OAuthFlow{Scopes: mockScopes}
	atURL := OAuthFlow{AuthorizationURL: exampleCom, TokenURL: exampleCom}
	arURL := OAuthFlow{AuthorizationURL: exampleCom, RefreshURL: exampleCom}
	aURLscopes := OAuthFlow{AuthorizationURL: exampleCom, Scopes: mockScopes}
	trURL := OAuthFlow{TokenURL: exampleCom, RefreshURL: exampleCom}
	tURLscopes := OAuthFlow{TokenURL: exampleCom, Scopes: mockScopes}
	rURLscopes := OAuthFlow{RefreshURL: exampleCom, Scopes: mockScopes}
	atrURL := OAuthFlow{AuthorizationURL: exampleCom, TokenURL: exampleCom, RefreshURL: exampleCom}
	atURLscopes := OAuthFlow{AuthorizationURL: exampleCom, TokenURL: exampleCom, Scopes: mockScopes}
	arURLscopes := OAuthFlow{AuthorizationURL: exampleCom, RefreshURL: exampleCom, Scopes: mockScopes}
	trURLscopes := OAuthFlow{TokenURL: exampleCom, RefreshURL: exampleCom, Scopes: mockScopes}
	atrURLscopes := OAuthFlow{AuthorizationURL: exampleCom, TokenURL: exampleCom, RefreshURL: exampleCom, Scopes: mockScopes}

	candidates := []struct {
		label   string
		in      OAuthFlow
		haveErr [4]bool
	}{
		{"empty", empty, [4]bool{true, true, true, true}},
		{"aURL", aURL, [4]bool{true, true, true, true}},
		{"tURL", tURL, [4]bool{true, true, true, true}},
		{"rURL", rURL, [4]bool{true, true, true, true}},
		{"scopes", scopes, [4]bool{true, true, true, true}},
		{"aURL/tURL", atURL, [4]bool{true, true, true, true}},
		{"aURL/rURL", arURL, [4]bool{true, true, true, true}},
		{"aURL/scopes", aURLscopes, [4]bool{false, true, true, true}},
		{"tURL/rURL", trURL, [4]bool{true, true, true, true}},
		{"tURL/scopes", tURLscopes, [4]bool{true, false, false, true}},
		{"rURL/scopes", rURLscopes, [4]bool{true, true, true, true}},
		{"aURL/tURL/rURL", atrURL, [4]bool{true, true, true, true}},
		{"aURL/tURL/scopes", atURLscopes, [4]bool{false, false, false, false}},
		{"aURL/rURL/scopes", arURLscopes, [4]bool{false, true, true, true}},
		{"tURL/rURL/scopes", trURLscopes, [4]bool{true, false, false, true}},
		{"aURL/tURL/rURL/scopes", atrURLscopes, [4]bool{false, false, false, false}},
	}
	for _, c := range candidates {
		testOAuthFlowValidate(t, c.label, c.in, c.haveErr)
	}
}

var flowTypes = []string{"implicit", "password", "clientCredentials", "authorizationCode"}

func testOAuthFlowValidate(t *testing.T, label string, oauthFlow OAuthFlow, haveErr [4]bool) {
	t.Logf("%s-empty", label)
	if err := oauthFlow.Validate(""); err == nil {
		t.Error("error should be occurred, but not")
	}
	t.Logf("%s-wrongtype", label)
	if err := oauthFlow.Validate("foobar"); err == nil {
		t.Error("error should be occurred, but not")
	}
	for i, flowType := range flowTypes {
		t.Logf("%s-%s", label, flowType)
		if err := oauthFlow.Validate(flowType); (err != nil) != haveErr[i] {
			if haveErr[i] {
				t.Error("error should be occurred, but not")
				continue
			}
			t.Error("error should not be occurred, but occurred")
			t.Log(err)
		}
	}
}
