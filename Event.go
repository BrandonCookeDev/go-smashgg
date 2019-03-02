package main

import (
	"github.com/tidwall/gjson"
)

type Event struct {
	id                     int64
	name                   string
	slug                   string
	state                  string
	startAt                int64
	numEntrants            int64
	checkInBuffer          int64
	checkInDuration        int64
	checkInEnabled         bool
	isOnline               bool
	teamNameAllowed        bool
	teamManagementDeadline int64
}

func GetEvent(tournamentSlug string, eventSlug string) Event {
	params := `{"slug":"tournament/` + tournamentSlug + `/event/` + eventSlug + `"}`
	data := query(eventQuery, params)
	return ParseEvent(data)
}

func ParseEvent(data string) Event {
	event := new(Event)
	event.id = gjson.Get(data, "data.event.id").Int()
	event.name = gjson.Get(data, "data.event.name").String()
	event.slug = gjson.Get(data, "data.event.slug").String()
	event.state = gjson.Get(data, "data.event.state").String()
	event.startAt = gjson.Get(data, "data.event.startAt").Int()
	event.numEntrants = gjson.Get(data, "data.event.numEntrants").Int()
	event.checkInBuffer = gjson.Get(data, "data.event.checkInBuffer").Int()
	event.checkInDuration = gjson.Get(data, "data.event.checkInDuration").Int()
	event.checkInEnabled = gjson.Get(data, "data.event.checkInEnabled").Bool()
	event.isOnline = gjson.Get(data, "data.event.isOnline").Bool()
	event.teamNameAllowed = gjson.Get(data, "data.event.teamNameAllowed").Bool()
	event.teamManagementDeadline = gjson.Get(data, "data.event.teamManagementDeadline").Int()
	return *event
}
