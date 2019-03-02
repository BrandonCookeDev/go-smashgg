package main

var tournamentQuery = `query TournamentQuery($slug: String){
	tournament(slug: $slug){
		id
		name
	}	
}`

var eventQuery = `query EventQuery($slug: String){
	event(slug: $slug){
		id
		name
	}	
}`