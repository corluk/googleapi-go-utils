package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetClientSecret(path string) (ClientSecret, error) {

	var clientSecret ClientSecret
	
	file, err := os.Open(path)
	if err != nil {
		return clientSecret, err

	}
	defer file.Close()

	bSecret, err := ioutil.ReadAll(file)
	if err != nil {
		return clientSecret, err

	}
	err = json.Unmarshal(bSecret, &clientSecret)

	if err != nil {
		return clientSecret, err

	}

	return clientSecret, nil

}

/*
func GetUrl(clientSecret ClientSecret, scopes []string, redirectUri *string) (string, error) {

	urlForm := new(AuthUrlForm)

	urlForm.MergeClient(clientSecret, scopes, redirectUri)

	urlForm.AccessType = "offline"
	urlForm.IncludeGrantedScopes = "true"
	urlForm.ResponseType = "code"
	urlForm.Bind()

	return urlForm.Url.String(), nil
}

func GetExchangeToken(token string, clientSecret ClientSecret, redirectUri *string) (TokenResponse, error) {

	var exchangeResponse TokenResponse
	var exchangeForm TokenForm
	if redirectUri == nil {
		redirectUri = &clientSecret.Installed.RedirectUris[0]
	}

	exchangeResponse, err := exchangeForm.Submit(token, &clientSecret, redirectUri)
	if err != nil {
		return exchangeResponse, err
	}

	return exchangeResponse, nil

}

func GetRefreshToken() {

}

func AddScope() {

}
*/
