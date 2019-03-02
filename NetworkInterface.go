package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var API_URL = "https://api.smash.gg/gql/alpha"

func query(query string, variables string) string {
	client := &http.Client{}
	authToken := os.Getenv("SMASHGG_API_TOKEN")

	message := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Cannot perform response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Cannot create response body")
	}

	return string(body) //json string
}
