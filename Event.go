package smashggo

import (
	"log"
	"strconv"

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

func GetEventSlug(tournamentSlug string, eventSlug string) string {
	return `tournament/` + tournamentSlug + `/event/` + eventSlug
}

func GetEvent(tournamentSlug string, eventSlug string) Event {
	params := `{"slug":` + GetEventSlug(tournamentSlug, eventSlug) + `}`
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

func GetEventSets(tournamentSlug string, eventSlug string, perPage int) []GGSet {
	var sets []GGSet

	page := 1
	params := `{
		"slug": "` + GetEventSlug(tournamentSlug, eventSlug) + `",
		"page": ` + strconv.Itoa(page) + `,
		"perPage": ` + strconv.Itoa(perPage) + `,
		"sortBy": null,
		"filters": null
	}`
	data := query(eventSetsQuery, params)
	phaseGroups := gjson.Get(data, "data.event.phaseGroups").Array()
	//log.Println(phaseGroups)

	var totalPages int64

	for _, pg := range phaseGroups {
		if totalPages == 0 {
			totalPages = gjson.Get(pg.String(), "paginatedSets.pageInfo.totalPages").Int()
		}

		pgSets := gjson.Get(pg.String(), "paginatedSets.nodes").Array()
		for _, set := range pgSets {
			ggSet := ParseGGSet(set.String())
			sets = append(sets, ggSet)
		}
	}

	//log.Printf("Got 1/%s Pages", string(totalPages))
	for i := 1; int64(i) <= totalPages; i++ {
		log.Printf("Got %d/%d Pages", i, totalPages)
		params = `{
			"slug": "` + GetEventSlug(tournamentSlug, eventSlug) + `",
			"page": ` + strconv.Itoa(page) + `,
			"perPage": ` + strconv.Itoa(perPage) + `,
			"sortBy": null,
			"filters": null
		}`

		data = query(eventSetsQuery, params)
		phaseGroups = gjson.Get(data, "data.event.phaseGroups").Array()

		for _, pg := range phaseGroups {
			if totalPages == 0 {
				totalPages = gjson.Get(pg.String(), "paginatedSets.pageInfo.totalPages").Int()
			}

			pgSets := gjson.Get(pg.String(), "paginatedSets.nodes").Array()
			for _, set := range pgSets {
				ggSet := ParseGGSet(set.String())
				sets = append(sets, ggSet)
			}
		}
	}

	//log.Println(sets)
	return sets
}
