package oas3

import "github.com/nasa9084/go-openapi/oauth"

// OAuthFlows Object
type OAuthFlows struct {
	Implicit          *OAuthFlow `yaml:",omitempty"`
	Password          *OAuthFlow `yaml:",omitempty"`
	ClientCredentials *OAuthFlow `yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode,omitempty"`
}

// Validate the values of OAuthFlows Object.
func (oauthFlows OAuthFlows) Validate() error {
	if oauthFlows.Implicit != nil {
		oauthFlows.Implicit.SetFlowType(oauth.ImplicitFlow)
		if err := oauthFlows.Implicit.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.Password != nil {
		oauthFlows.Password.SetFlowType(oauth.PasswordFlow)
		if err := oauthFlows.Password.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.ClientCredentials != nil {
		oauthFlows.ClientCredentials.SetFlowType(oauth.ClientCredentialsFlow)
		if err := oauthFlows.ClientCredentials.Validate(); err != nil {
			return err
		}
	}
	if oauthFlows.AuthorizationCode != nil {
		oauthFlows.AuthorizationCode.SetFlowType(oauth.AuthorizationCodeFlow)
		if err := oauthFlows.AuthorizationCode.Validate(); err != nil {
			return err
		}
	}
	return nil
}
