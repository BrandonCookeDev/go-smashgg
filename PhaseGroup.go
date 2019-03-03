package smashggo

import (
	"log"
	"strconv"

	"github.com/tidwall/gjson"
)

type PhaseGroup struct {
	id                int64
	displayIdentifier string
	firstRoundTime    int64
	state             int64
	phaseId           int64
	waveId            int64
}

func GetPhaseGroup(id int) PhaseGroup {
	params := `{"id":` + strconv.Itoa(id) + `}`
	data := query(phaseGroupQuery, params)
	return ParsePhaseGroup(data)
}

func ParsePhaseGroup(data string) PhaseGroup {
	phaseGroup := new(PhaseGroup)
	phaseGroup.id = gjson.Get(data, "data.phaseGroup.id").Int()
	phaseGroup.displayIdentifier = gjson.Get(data, "data.phaseGroup.displayIdentifier").String()
	phaseGroup.firstRoundTime = gjson.Get(data, "data.phaseGroup.firstRoundTime").Int()
	phaseGroup.state = gjson.Get(data, "data.phaseGroup.state").Int()
	phaseGroup.phaseId = gjson.Get(data, "data.phaseGroup.phaseId").Int()
	phaseGroup.waveId = gjson.Get(data, "data.phaseGroup.waveId").Int()
	return *phaseGroup
}

func GetPhaseGroupSets(id int, perPage int) []GGSet {
	var sets []GGSet
	var totalPages int64

	page := 1
	params := `{
		"id": ` + strconv.Itoa(id) + `,
		"page": ` + strconv.Itoa(page) + `,
		"perPage": ` + strconv.Itoa(perPage) + `,
		"sortType": null,
		"filters": null
	}`

	data := query(phaseGroupSetsQuery, params)
	totalPages = gjson.Get(data, "phaseGroup.paginatedSets.pageInfo.totalPages").Int()
	pgSets := gjson.Get(data, "phaseGroup.paginatedSets.nodes").Array()

	for _, set := range pgSets {
		ggSet := ParseGGSet(set.String())
		sets = append(sets, ggSet)
	}

	//log.Printf("Got 1/%s Pages", string(totalPages))
	for i := 1; int64(i) <= totalPages; i++ {
		log.Printf("Got %d/%d Pages", i, totalPages)
		params = `{
			"id": ` + strconv.Itoa(id) + `,
			"page": ` + strconv.Itoa(page) + `,
			"perPage": ` + strconv.Itoa(perPage) + `,
			"sortType": null,
			"filters": null
		}`

		data = query(phaseGroupSetsQuery, params)
		totalPages = gjson.Get(data, "phaseGroup.paginatedSets.pageInfo.totalPages").Int()
		pgSets = gjson.Get(data, "phaseGroup.paginatedSets.nodes").Array()

		for _, set := range pgSets {
			ggSet := ParseGGSet(set.String())
			sets = append(sets, ggSet)
		}
	}

	//log.Println(sets)
	return sets
}
