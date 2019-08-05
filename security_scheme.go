package oas3

// SecurityScheme Object
type SecurityScheme struct {
	Type             SecuritySchemeType `yaml:",omitempty"`
	Description      string             `yaml:",omitempty"`
	Name             string             `yaml:",omitempty"`
	In               InType             `yaml:",omitempty"`
	Scheme           string             `yaml:",omitempty"`
	BearerFormat     string             `yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows        `yaml:",omitempty"`
	OpenIDConnectURL string             `yaml:"openIdConnectUrl,omitempty"`

	Ref string `yaml:"$ref,omitempty"`
}

// SecuritySchemeType represents a securityScheme.type value.
type SecuritySchemeType string

// SecuritySchemeTypes
const (
	APIKeyType        SecuritySchemeType = "apiKey"
	HTTPType          SecuritySchemeType = "http"
	OAuth2Type        SecuritySchemeType = "oauth2"
	OpenIDConnectType SecuritySchemeType = "openIdConnect"
)

// SecuritySchemeTypeList is a list of valid values of securityScheme.Type.
var SecuritySchemeTypeList = []string{string(APIKeyType), string(HTTPType), string(OAuth2Type), string(OpenIDConnectType)}

// Validate the values of SecurityScheme object.
func (secScheme SecurityScheme) Validate() error {
	switch secScheme.Type {
	case "":
		return ErrRequired{Target: "securityScheme.type"}
	case APIKeyType:
		return secScheme.validateFieldForAPIKey()
	case HTTPType:
		return secScheme.validateFieldForHTTP()
	case OAuth2Type:
		return secScheme.validateFieldForOAuth2()
	case OpenIDConnectType:
		return secScheme.validateFieldForOpenIDConnect()
	}
	return ErrMustOneOf{Object: "securityScheme.type", ValidValues: SecuritySchemeTypeList}
}

func (secScheme SecurityScheme) validateFieldForAPIKey() error {
	if secScheme.Name == "" {
		return ErrRequired{"securityScheme.name"}
	}
	if secScheme.In == "" {
		return ErrRequired{"securityScheme.in"}
	}
	if secScheme.In != InQuery && secScheme.In != InHeader && secScheme.In != InCookie {
		return ErrMustOneOf{Object: "securityScheme.in", ValidValues: SecuritySchemeInList}
	}
	return nil
}

func (secScheme SecurityScheme) validateFieldForHTTP() error {
	if secScheme.Scheme == "" {
		return ErrRequired{Target: "securityScheme.scheme"}
	}
	return nil
}

func (secScheme SecurityScheme) validateFieldForOAuth2() error {
	if secScheme.Flows == nil {
		return ErrRequired{Target: "securityScheme.flows"}
	}
	return secScheme.Flows.Validate()
}

func (secScheme SecurityScheme) validateFieldForOpenIDConnect() error {
	return mustURL("securityScheme.openIdConnectUrl", secScheme.OpenIDConnectURL)
}
