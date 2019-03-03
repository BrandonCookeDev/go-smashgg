package gosmashgg

import (
	"log"
	"regexp"
	"strconv"

	"github.com/tidwall/gjson"
)

var displayScoreRegex, _ = regexp.Compile("([\\S\\s]*) ([0-9]{1,3}) - ([\\S\\s]*) ([0-9]{1,3})")

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
	player1       string
	player2       string
	player1Score  int
	player2Score  int
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
	parsed := set.ParseDisplayScore()
	if len(parsed) > 0 {
		set.player1 = parsed[1]
		set.player1Score, _ = strconv.Atoi(parsed[2])
		set.player2 = parsed[3]
		set.player2Score, _ = strconv.Atoi(parsed[4])
	}
	return *set
}

func (set *GGSet) ParseDisplayScore() []string {
	return displayScoreRegex.FindStringSubmatch(set.displayScore)
}

func (set *GGSet) Print() {
	log.Printf("%s: %s %d - %d %s", set.fullRoundText, set.player1, set.player1Score, set.player2Score, set.player2)
}
