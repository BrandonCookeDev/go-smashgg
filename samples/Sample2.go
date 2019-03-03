package main

import (
	"log"
	"smashggo"
)

func main() {
	sets := smashggo.GetEventSets("function-1-recursion-regional", "melee-singles", 5)
	log.Println(sets)
}
