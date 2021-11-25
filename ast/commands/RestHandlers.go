package commands

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetHandler(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if messageBody, err := ioutil.ReadAll(resp.Body); err == nil {

		dataPackage, err := CreatePackageFromBytes(messageBody)

		if err != nil {
			log.Fatal().Err(err)
			return nil, err
		} else {
			return dataPackage.FieldMap, nil
		}
	} else {
		log.Fatal().Err(err)
		return nil, err
	}
}

func PostHandler(url string, body map[string]interface{}) error {

	messageBody, _ := json.Marshal(body)
	encodedBody := bytes.NewBuffer(messageBody)

	response, err := http.Post(url, "application/json", encodedBody)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	stringBody := string(responseBody)
	log.Info().Msgf("Body: %s", stringBody)

	return nil
}
