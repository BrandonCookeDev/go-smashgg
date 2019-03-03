package smashggo

import "github.com/tidwall/gjson"

type GGSet struct {
	id            int64
	eventId       int64
	phaseGroupId  int64
	displayScore  string
	fullRoundText string
	round         string
	startedAt     int64
	completedAt   int64
	winnerId      int64
	totalGames    int64
	state         int64
}

func ParseGGSet(data string) GGSet {
	set := new(GGSet)
	set.id = gjson.Get(data, "id").Int()
	set.eventId = gjson.Get(data, "eventId").Int()
	set.phaseGroupId = gjson.Get(data, "phaseGroupId").Int()
	set.displayScore = gjson.Get(data, "displayScore").String()
	set.fullRoundText = gjson.Get(data, "fullRoundText").String()
	set.round = gjson.Get(data, "round").String()
	set.startedAt = gjson.Get(data, "startedAt").Int()
	set.completedAt = gjson.Get(data, "completedAt").Int()
	set.winnerId = gjson.Get(data, "winnerId").Int()
	set.totalGames = gjson.Get(data, "totalGames").Int()
	set.state = gjson.Get(data, "state").Int()
	return *set
}
