package auth

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func setPrompt(loginUrl RequestLoginUrl, q url.Values) error {

	if len(loginUrl.Prompt) > 0 {

		for _, prompt := range loginUrl.Prompt {

			switch prompt {
			case "none":
			case "consent":
			case "select_account":
				break
			default:
				return errors.New("unknown prompt " + prompt)
				break
			}
		}
		q.Add("prompt", strings.Join(loginUrl.Prompt, " "))
	}
	return nil
}
func setLoginHint(loginUrl RequestLoginUrl, q url.Values) {

	if strings.TrimSpace(loginUrl.LoginHint) != "" {
		q.Add("login_hint", loginUrl.LoginHint)
	}

}
func setAccessType(loginUrl RequestLoginUrl, q url.Values) error {

	var accessType string = "online"
	switch loginUrl.AccessType {
	case "offline":
	case "online":
		accessType = loginUrl.AccessType
		break
	default:
		return errors.New("wrong type of access type set please  select online or offline")
		break
	}
	q.Add("access_type", accessType)
	return nil

}

func setIncludeGrantedScopes(loginUrl RequestLoginUrl, q url.Values) {
	var value bool = true
	if loginUrl.IncludeGrantedScopes != nil {

		value = *loginUrl.IncludeGrantedScopes

	}
	q.Add("include_granted_scopes", strconv.FormatBool(value))
}
func setState(loginUrl RequestLoginUrl, q url.Values) {

	if strings.TrimSpace(loginUrl.State) != "" {
		q.Add("state", loginUrl.State)
	}
}
func setScope(loginUrl RequestLoginUrl, q url.Values) error {

	if len(loginUrl.Scope) < 1 {
		return errors.New("You must set at least one Scope")
	}
	q.Add("scope", strings.Join(loginUrl.Scope, " "))
	return nil

}
func setRedirectUri(loginUrl RequestLoginUrl, q url.Values) error {

	if strings.TrimSpace(loginUrl.RedirectURI) == "" {
		return errors.New("You must set redirect URI")
	}
	q.Add("redirect_uri", loginUrl.RedirectURI)
	return nil
}
func setClientId(loginUrl RequestLoginUrl, q url.Values) error {

	if strings.TrimSpace(loginUrl.ClientId) == "" {
		return errors.New("ClientId must set")
	}
	q.Add("client_id", loginUrl.ClientId)
	return nil
}
func setResponseType(loginUrl RequestLoginUrl, q url.Values) {
	var responseType string = "code"
	switch loginUrl.ResponseType {

	case "code":
		responseType = "code"

	}
	q.Add("response_type", responseType)

}
func (loginUrl RequestLoginUrl) GetUrl() (*url.URL, error) {

	uri, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")

	if err != nil {
		return uri, err
	}

	q := uri.Query()

	err = setClientId(loginUrl, q)
	if err != nil {
		return uri, err
	}

	err = setRedirectUri(loginUrl, q)
	if err != nil {
		return uri, err
	}
	setResponseType(loginUrl, q)

	err = setScope(loginUrl, q)
	if err != nil {
		return uri, err
	}

	err = setAccessType(loginUrl, q)
	if err != nil {
		return uri, err
	}
	setState(loginUrl, q)
	setIncludeGrantedScopes(loginUrl, q)
	setLoginHint(loginUrl, q)
	err = setPrompt(loginUrl, q)
	if err != nil {
		return uri, err
	}

	uri.RawQuery = q.Encode()
	return uri, nil
}

func OpenBrowser(url string) error {

	fmt.Println(url)
	switch runtime.GOOS {
	case "linux":
		err := exec.Command("xdg-open", url).Start()
		if err != nil {
			return err
		}
	case "windows":
		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		if err != nil {
			return err
		}
	case "darwin":
		err := exec.Command("open", url).Start()
		if err != nil {
			return err
		}
	default:
		return errors.New("cannot detect browser")
	}

	return nil

}
