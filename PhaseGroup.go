package main

import (
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
