package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "io/ioutil"
	"os"
	"github.com/tidwall/gjson"
)

var API_URL = "https://api.smash.gg/gql/alpha"

func main() {
	fmt.Printf("API URL is: %s", API_URL)

	authToken := os.Getenv("SMASHGG_API_TOKEN")
	log.Printf("API TOKEN set to %s", authToken)

	to12 := GetTournament(authToken, "to12")
	ceo := GetTournament(authToken, "ceo2017")

	log.Println(to12)
	log.Println(ceo)
}

func GetTournament(authToken string, slug string) Tournament {
	client := &http.Client{}

	message := map[string]interface{}{
		"query": tournamentQuery,
		"variables": "{\"slug\": \"" + slug + "\"}",
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
	if err != nil{
		log.Fatalln("Cannot perform response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln("Cannot create response body")
	}

	var resultJson map[string]interface{}
	json.Unmarshal(body, &resultJson)

	tournament := new(Tournament)
	tournament.id = gjson.Get(string(body), "data.tournament.id").Int()
	tournament.name = gjson.Get(string(body), "data.tournament.name").String()
	

	return *tournament
}
