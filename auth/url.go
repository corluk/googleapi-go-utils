package auth

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

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

func (authFlow AuthFlow) GetUrl() (string, error) {

	if authFlow.AuthUrlForm == nil {
		authFlow.AuthUrlForm = new(AuthUrlForm)
	}

	uri, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")

	if err != nil {
		return "", err
	}

	q := uri.Query()

	responseType := authFlow.AuthUrlForm.ResponseType
	if responseType == nil {
		code := "code"
		responseType = &code
	}
	redirectUri := authFlow.RedirectUri
	if redirectUri == nil {
		redirectUri = &authFlow.ClientSecret.Installed.RedirectUris[0]
	}
	scopes := authFlow.Scopes
	if scopes == nil {
		return "", errors.New("you must set scopes before requesting url ")
	}

	q.Add("client_id", authFlow.ClientSecret.Installed.ClientId)
	q.Add("redirect_uri", *redirectUri)

	q.Add("scope", strings.Join(*scopes, ","))
	q.Add("access_type", authFlow.AuthUrlForm.AccessType)
	q.Add("include_granted_scopes", authFlow.AuthUrlForm.IncludeGrantedScopes)
	q.Add("response_type", *responseType)

	uri.RawQuery = q.Encode()
	return uri.String(), nil
}

func (authFlow *AuthFlow) OpenBrowser() error {

	url, err := authFlow.GetUrl()
	if err != nil {
		return err
	}
	fmt.Println(url)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
		if err != nil {
			return err
		}
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		if err != nil {
			return err
		}
	case "darwin":
		err = exec.Command("open", url).Start()
		if err != nil {
			return err
		}
	default:
		return errors.New("cannot detect browser")
	}
	if err != nil {
		return err
	}

	return nil

}
