package oas3

import (
	"net/url"

	"github.com/nasa9084/go-openapi/oauth"
)

// OAuthFlow Object
type OAuthFlow struct {
	flowType         string            `yaml:",omitempty"`
	AuthorizationURL string            `yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `yaml:",omitempty"`
}

var defined = struct{}{}

var validFlowTypes = map[string]struct{}{
	oauth.ImplicitFlow:          defined,
	oauth.PasswordFlow:          defined,
	oauth.ClientCredentialsFlow: defined,
	oauth.AuthorizationCodeFlow: defined,
}

var requireAuthorizationURL = map[string]struct{}{
	oauth.ImplicitFlow:          defined,
	oauth.AuthorizationCodeFlow: defined,
}

var requireTokenURL = map[string]struct{}{
	oauth.PasswordFlow:          defined,
	oauth.ClientCredentialsFlow: defined,
	oauth.AuthorizationCodeFlow: defined,
}

// SetFlowType sets oauth flow type.
func (oauthFlow *OAuthFlow) SetFlowType(typ string) {
	oauthFlow.flowType = typ
}

// Validate the values of OAuthFlow object.
func (oauthFlow OAuthFlow) Validate() error {
	if _, ok := validFlowTypes[oauthFlow.flowType]; !ok {
		return ErrInvalidFlowType
	}
	if _, ok := requireAuthorizationURL[oauthFlow.flowType]; ok {
		if err := mustURL("oauthFlow.authorizationUrl", oauthFlow.AuthorizationURL); err != nil {
			return err
		}
	}
	if _, ok := requireTokenURL[oauthFlow.flowType]; ok {
		if err := mustURL("oauthFlow.tokenUrl", oauthFlow.TokenURL); err != nil {
			return err
		}
	}
	if oauthFlow.RefreshURL != "" {
		if _, err := url.ParseRequestURI(oauthFlow.RefreshURL); err != nil {
			return ErrFormatInvalid{Target: "oauthFlow.refreshUrl", Format: "URL"}
		}
	}
	if oauthFlow.Scopes == nil || len(oauthFlow.Scopes) == 0 {
		return ErrRequired{Target: "oauthFlow.scopes"}
	}

	return nil
}
