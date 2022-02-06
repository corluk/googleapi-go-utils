package auth

import "net/url"

type RefreshForm struct {
	ClientId     string
	ClientSecret string
	RefreshToken string
	GrantType    string
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Credentials struct {
	ClientId  string `json:"client_id"`
	ProjectId string `json:"project_id"`

	AuthUri                 string   `json:"auth_uri"`
	TokenUri                string   `json:"token_uri"`
	AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
	RedirectUris            []string `json:"redirect_uris"`

	ClientSecret string `json:"client_secret"`
	ApiKey       string `json:"apiKey"`
}
type ClientSecret struct {
	Installed Credentials `json:"installed"`
}

type AuthUrlForm struct {
	ClientId             string
	RedirectUri          *string
	Scopes               []string
	AccessType           string
	IncludeGrantedScopes string
	ResponseType         *string
	Url                  *url.URL
}

type TokenForm struct {
	Code         string  `json:"code"`
	ClientId     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	RedirectUri  *string `json:"redirect_uri"`
	GrantType    string  `json:"grant_type"`
}

type TokenResponse struct {
	ID           string `json:"_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
}

type AuthFlow struct {
	ExchangeForm     *TokenForm
	ExchangeResponse *TokenResponse
	ClientSecret     *ClientSecret
	Scopes           *[]string
	RedirectUri      *string
	AuthUrlForm      *AuthUrlForm
	RefreshForm      *RefreshForm
}

type RequestLoginUrl struct {
	ClientId             string   `json:"client_id"`
	RedirectURI          string   `json:"redirect_uri"`
	ResponseType         string   `json:"response_type"`
	Scope                []string `json:"scope"`
	AccessType           string   `json:"access_type"`
	State                string   `json:"state"`
	IncludeGrantedScopes *bool    `json:"include_granted_scopes"`
	LoginHint            string   `json:"login_hint"`
	Prompt               []string `json:"prompt"`
}
