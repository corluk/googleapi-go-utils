package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)


func (authFlow *AuthFlow) SaveToken(path string) error {

	bExchangeResponse, err := json.Marshal(authFlow.ExchangeResponse)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, bExchangeResponse, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (authFlow *AuthFlow) RefreshToken() error {
	endPoint := "https://oauth2.googleapis.com/token"
	formData := url.Values{}
	formData.Set("client_id", authFlow.ClientSecret.Installed.ClientId)
	formData.Set("client_secret", authFlow.ClientSecret.Installed.ClientSecret)
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", authFlow.ExchangeResponse.RefreshToken)
	request, err := http.NewRequest("POST", endPoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("status code %d", response.StatusCode)
		return errors.New("unexpected status code" + string(response.StatusCode))
	}
	bResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	err = json.Unmarshal(bResponse, &authFlow.ExchangeResponse)
	if err != nil {
		return err
	}
	return nil

}
func (authFlow *AuthFlow) ExchangeToken(token string) error {

	if authFlow.ClientSecret == nil {
		return errors.New("you must set client secret before requesting access token ")
	}

	if authFlow.ExchangeForm == nil {
		authFlow.ExchangeForm = new(TokenForm)
	}
	redirectUri := authFlow.RedirectUri
	if redirectUri == nil {
		redirectUri = &authFlow.ClientSecret.Installed.RedirectUris[0]
	}
	authFlow.ExchangeForm.Code = token
	authFlow.ExchangeForm.ClientId = authFlow.ClientSecret.Installed.ClientId
	authFlow.ExchangeForm.ClientSecret = authFlow.ClientSecret.Installed.ClientSecret
	authFlow.ExchangeForm.GrantType = "authorization_code"
	authFlow.ExchangeForm.RedirectUri = redirectUri

	formData := url.Values{}
	formData.Set("client_id", authFlow.ExchangeForm.ClientId)
	formData.Set("client_secret", authFlow.ExchangeForm.ClientSecret)
	formData.Set("code", authFlow.ExchangeForm.Code)
	formData.Set("redirectUri", *authFlow.ExchangeForm.RedirectUri)
	formData.Set("grant_type", authFlow.ExchangeForm.GrantType)
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

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected status code" + string(response.StatusCode))
	}
	bResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	err = json.Unmarshal(bResponse, &authFlow.ExchangeResponse)
	if err != nil {
		return err
	}
	if len(authFlow.ExchangeResponse.AccessToken) < 1 {
		return errors.New(string(bResponse))
	}
	return nil
}

func (authFlow *AuthFlow) LoadToken(path string) error {
	fmt.Println("File reading")
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println(file)
	bFile, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}
	fmt.Println(string(bFile))
	err = json.Unmarshal(bFile, &authFlow.ExchangeResponse)

	if err != nil {
		return err
	}

	return nil

}

func (authFlow *AuthFlow) LoadSecretFile(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err

	}
	defer file.Close()

	bSecret, err := ioutil.ReadAll(file)
	if err != nil {
		return err

	}
	err = json.Unmarshal(bSecret, &authFlow.ClientSecret)

	if err != nil {
		return err

	}

	return nil

}
func (authFlow *AuthFlow) Init(pathClientSecret string, pathToken string) error {
	err := authFlow.LoadSecretFile(pathClientSecret)
	if err != nil {
		return err
	}
	err = authFlow.LoadToken(pathToken)
	if err != nil {
		return err
	}
	return nil
}
func (authFlow *AuthFlow) Request(request *http.Request, object interface{}) error {

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authFlow.ExchangeResponse.AccessToken))
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	fmt.Printf("out going url %s", request.URL.Query().Encode())
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected status code" + fmt.Sprint(response.StatusCode))
	}

	err = json.Unmarshal(bResponse, &object)
	if err != nil {
		return err
	}
	return nil
}
func (authFlow *AuthFlow) PostForm(endPoint *url.URL, formData *url.Values, object interface{}) error {

	request, err := http.NewRequest("POST", endPoint.String(), strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return authFlow.Request(request, &object)
}
func (authFlow *AuthFlow) Get(endPoint *url.URL, object interface{}) error {

	request, err := http.NewRequest("GET", endPoint.String(), nil)
	if err != nil {
		return err
	}

	return authFlow.Request(request, &object)

}
