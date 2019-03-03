package main

import (
	"log"
	"os"
	"smashggo"
)

func main2() {

	authToken := os.Getenv("SMASHGG_API_TOKEN")
	log.Printf("API TOKEN set to %s", authToken)

	to12 := smashggo.GetTournament("to12")
	ceo := smashggo.GetTournament("ceo2017")

	log.Println(to12)
	log.Println(ceo)

	to12Melee := smashggo.GetEvent("ceo-2016", "melee-singles")
	ceoMelee := smashggo.GetEvent("function-1-recursion-regional", "melee-singles")

	log.Println(to12Melee)
	log.Println(ceoMelee)

	to12PhaseGroup1 := smashggo.GetPhaseGroup(301994)
	to12PhaseGroup2 := smashggo.GetPhaseGroup(373938)

	log.Println(to12PhaseGroup1)
	log.Println(to12PhaseGroup2)

	to12Phase1 := smashggo.GetPhase(100046)
	to12Phase2 := smashggo.GetPhase(132397)

	log.Println(to12Phase1)
	log.Println(to12Phase2)
}
