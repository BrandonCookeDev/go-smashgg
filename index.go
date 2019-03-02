package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var API_URL = "https://api.smash.gg/gql/alpha"

func main() {
	fmt.Printf("API URL is: %s", API_URL)
	GetTournament("to12")
	GetTournament("ceo2017")
}

func GetTournament(slug string) {
	client := &http.Client{}

	authToken := os.Getenv("SMASHGG_API_TOKEN")
	log.Printf("API TOKEN set to %s", authToken)

	message := map[string]interface{}{
		"query": `query TournamentQuery($slug: String){
			tournament(slug: $slug){
				id
				name
				events{
					id
					name
				}
			}	
		}`,
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

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
