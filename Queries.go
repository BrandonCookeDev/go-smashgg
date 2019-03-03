package smashggo

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
		slug
		state
		startAt
		numEntrants
		checkInBuffer
		checkInDuration
		checkInEnabled
		isOnline
		teamNameAllowed
		teamManagementDeadline
	}	
}`

var phaseQuery = `query PhaseQuery($id: Int){
	phase(id: $id){
		id
		name
		numSeeds
		groupCount
	}	
}`

var phaseGroupQuery = `query PhaseGroupQuery($id: Int){
	phaseGroup(id: $id){
		id
		displayIdentifier
		firstRoundTime
		state
		phaseId
		waveId
		tiebreakOrder
	}	
}`

var eventSetsQuery = `query EventSets($slug: String, $page: Int, $perPage: Int, $sortType: SetSortType, $filters: SetFilters){
	event(slug: $slug){
		phaseGroups{
			paginatedSets(
				page: $page,
				perPage: $perPage,
				sortType: $sortType,
				filters: $filters
			){
				pageInfo{
					totalPages
				}
				nodes{
					id
					eventId
					phaseGroupId
					displayScore
					fullRoundText
					round
					startedAt
					completedAt
					winnerId
					totalGames
					state
				}
			}
		}
	}	
}`

var phaseGroupSetsQuery = `query PhaseGroupSets($id: Int, $page: Int, $perPage: Int, $sortType: SetSortType, $filters: SetFilters){
	phaseGroup(id: $id){
		paginatedSets(
			page: $page,
			perPage: $perPage,
			sortType: $sortType,
			filters: $filters
		){
			pageInfo{
				totalPages
			}
			nodes{
				id
				eventId
				phaseGroupId
				displayScore
				fullRoundText
				round
				startedAt
				completedAt
				winnerId
				totalGames
				state
			}
		}
	}	
}`
