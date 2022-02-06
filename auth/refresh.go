package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (authFlow *AuthFlow) SubmitRefreshForm(clientSecret ClientSecret, exchangeResponse TokenResponse) error {

	refreshForm := new(RefreshForm)
	if authFlow.ClientSecret == nil || authFlow.ExchangeResponse == nil {
		return errors.New("either clientSecret or ExchangeResponse are nil ")
	}
	refreshForm.ClientId = authFlow.ClientSecret.Installed.ClientId
	refreshForm.ClientSecret = authFlow.ClientSecret.Installed.ClientSecret
	refreshForm.RefreshToken = authFlow.ExchangeResponse.RefreshToken

	formData := url.Values{}
	formData.Set("client_id", refreshForm.ClientId)
	formData.Set("client_secret", refreshForm.ClientSecret)
	formData.Set("refreshToken", refreshForm.RefreshToken)
	formData.Set("grant_type", "refresh_token")

	endPoint := "https://oauth2.googleapis.com/token"
	request, err := http.NewRequest("POST", endPoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	bExchangeResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bExchangeResponse, &authFlow.ExchangeResponse)
	if err != nil {
		return err
	}

	return nil

}
