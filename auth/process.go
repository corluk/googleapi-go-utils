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
