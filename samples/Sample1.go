package main

import (
	"gosmashgg"
	"log"
	"os"
)

func main2() {

	authToken := os.Getenv("SMASHGG_API_TOKEN")
	log.Printf("API TOKEN set to %s", authToken)

	to12 := gosmashgg.GetTournament("to12")
	ceo := gosmashgg.GetTournament("ceo2017")

	log.Println(to12)
	log.Println(ceo)

	to12Melee := gosmashgg.GetEvent("ceo-2016", "melee-singles")
	ceoMelee := gosmashgg.GetEvent("function-1-recursion-regional", "melee-singles")

	log.Println(to12Melee)
	log.Println(ceoMelee)

	to12PhaseGroup1 := gosmashgg.GetPhaseGroup(301994)
	to12PhaseGroup2 := gosmashgg.GetPhaseGroup(373938)

	log.Println(to12PhaseGroup1)
	log.Println(to12PhaseGroup2)

	to12Phase1 := gosmashgg.GetPhase(100046)
	to12Phase2 := gosmashgg.GetPhase(132397)

	log.Println(to12Phase1)
	log.Println(to12Phase2)
}
