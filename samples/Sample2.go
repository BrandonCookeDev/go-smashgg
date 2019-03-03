package main

import (
	"gosmashgg"
	"log"
)

func main() {
	sets := gosmashgg.GetEventSets("function-1-recursion-regional", "melee-singles", 5)
	log.Println(sets)

	for _, set := range sets {
		set.Print()
	}

	sets2 := gosmashgg.GetPhaseGroupSets(301994, 5)
	log.Println(sets2)
}
