package oas3_test

import (
	"testing"

	openapi "github.com/crolly/oas3"
)

func TestOAuthFlows_Validate(t *testing.T) {
	authorizationURLRequiredError := openapi.ErrRequired{Target: "oauthFlow.authorizationUrl"}
	tokenURLRequiredError := openapi.ErrRequired{Target: "oauthFlow.tokenUrl"}
	candidates := []candidate{
		{"empty", openapi.OAuthFlows{}, nil},
		{"invalidImplicit", openapi.OAuthFlows{Implicit: &openapi.OAuthFlow{}}, authorizationURLRequiredError},
		{"invalidPassword", openapi.OAuthFlows{Password: &openapi.OAuthFlow{}}, tokenURLRequiredError},
		{"invalidClientCredentials", openapi.OAuthFlows{ClientCredentials: &openapi.OAuthFlow{}}, tokenURLRequiredError},
		{"invalidAuthorizationCode", openapi.OAuthFlows{AuthorizationCode: &openapi.OAuthFlow{}}, authorizationURLRequiredError},
	}
	testValidater(t, candidates)
}
