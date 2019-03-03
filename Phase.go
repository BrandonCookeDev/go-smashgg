package smashggo

import (
	"strconv"

	"github.com/tidwall/gjson"
)

type Phase struct {
	id         int64
	name       string
	numSeeds   int64
	groupCount int64
}

func GetPhase(id int) Phase {
	params := `{"id":` + strconv.Itoa(id) + `}`
	data := query(phaseQuery, params)
	return ParsePhase(data)
}

func ParsePhase(data string) Phase {
	phase := new(Phase)
	phase.id = gjson.Get(data, "data.phase.id").Int()
	phase.name = gjson.Get(data, "data.phase.name").String()
	phase.numSeeds = gjson.Get(data, "data.phase.numSeeds").Int()
	phase.groupCount = gjson.Get(data, "data.phase.groupCount").Int()
	return *phase
}
