package main

import "github.com/tidwall/gjson"

type Event struct {
	id   int64
	name string
}

func GetEvent(tournamentSlug string, eventSlug string) Event {
	params := `{"slug":"tournament/` + tournamentSlug + `/event/` + eventSlug + `"}`
	data := query(eventQuery, params)
	event := new(Event)
	event.id = gjson.Get(data, "data.event.id").Int()
	event.name = gjson.Get(data, "data.event.name").String()
	return *event
}
