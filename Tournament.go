package gosmashgg

import (
	"github.com/tidwall/gjson"
)

type Tournament struct {
	id             int64
	name           string
	slug           string
	city           string
	postalCode     string
	addrState      string
	countryCode    string
	region         string
	venueAddress   string
	venueName      string
	gettingThere   string
	lat            float64
	lng            float64
	timezone       string
	startAt        int64
	endAt          int64
	contactInfo    string
	contactEmail   string
	contactTwitter string
	contactPhone   string
	ownerId        string
}

func GetTournament(slug string) Tournament {
	params := `{"slug":"` + slug + `"}`
	data := query(tournamentQuery, params)
	return ParseTournament(data)
}

func ParseTournament(data string) Tournament {
	tournament := new(Tournament)
	tournament.id = gjson.Get(data, "data.tournament.id").Int()
	tournament.name = gjson.Get(data, "data.tournament.name").String()
	tournament.slug = gjson.Get(data, "data.tournament.slug").String()
	tournament.city = gjson.Get(data, "data.tournament.city").String()
	tournament.postalCode = gjson.Get(data, "data.tournament.postalCode").String()
	tournament.addrState = gjson.Get(data, "data.tournament.addrState").String()
	tournament.countryCode = gjson.Get(data, "data.tournament.countryCode").String()
	tournament.region = gjson.Get(data, "data.tournament.region").String()
	tournament.venueAddress = gjson.Get(data, "data.tournament.venueAddress").String()
	tournament.venueName = gjson.Get(data, "data.tournament.venueName").String()
	tournament.gettingThere = gjson.Get(data, "data.tournament.gettingThere").String()
	tournament.lat = gjson.Get(data, "data.tournament.lat").Float()
	tournament.lng = gjson.Get(data, "data.tournament.lng").Float()
	tournament.timezone = gjson.Get(data, "data.tournament.timezone").String()
	tournament.startAt = gjson.Get(data, "data.tournament.startAt").Int()
	tournament.endAt = gjson.Get(data, "data.tourament.endAt").Int()
	tournament.contactInfo = gjson.Get(data, "data.tournament.contactInfo").String()
	tournament.contactEmail = gjson.Get(data, "data.tournament.contactEmail").String()
	tournament.contactPhone = gjson.Get(data, "data.tournament.contactPhone").String()
	tournament.contactTwitter = gjson.Get(data, "data.tournament.contactTwitter").String()
	return *tournament
}
