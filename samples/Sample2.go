package main

import (
	"log"
	"smashggo"
)

func main() {
	sets := smashggo.GetEventSets("function-1-recursion-regional", "melee-singles", 5)
	log.Println(sets)

	sets2 := smashggo.GetPhaseGroupSets(301994, 5)
	log.Println(sets2)
}
