package main

var tournamentQuery = `query TournamentQuery($slug: String){
	tournament(slug: $slug){
		id
		name
		slug
		city
		postalCode
		addrState
		countryCode
		region
		venueAddress
		venueName
		gettingThere
		lat
		lng
		timezone
		startAt
		endAt
		contactInfo
		contactEmail
		contactTwitter
		contactPhone
		ownerId
	}	
}`

var eventQuery = `query EventQuery($slug: String){
	event(slug: $slug){
		id
		name
	}	
}`
