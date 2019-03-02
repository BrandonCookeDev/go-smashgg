package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("API URL is: %s", API_URL)

	authToken := os.Getenv("SMASHGG_API_TOKEN")
	log.Printf("API TOKEN set to %s", authToken)

	to12 := GetTournament("to12")
	ceo := GetTournament("ceo2017")

	log.Println(to12)
	log.Println(ceo)

	to12Melee := GetEvent("ceo-2016", "melee-singles")
	ceoMelee := GetEvent("function-1-recursion-regional", "melee-singles")

	log.Println(to12Melee)
	log.Println(ceoMelee)

	to12PhaseGroup1 := GetPhaseGroup(301994)
	to12PhaseGroup2 := GetPhaseGroup(373938)

	log.Println(to12PhaseGroup1)
	log.Println(to12PhaseGroup2)
}
