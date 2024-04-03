package main

import("sort")

func sortTeams(players []Player, size int) []Team {
	if len(players) % size != 0 {
		panic("Team size results in unassigned players.")
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Rating > players[j].Rating
	})
	
	var teams = make([]Team, size)

	for i := 0; i < len(players) / size ; i++ {
		for j := 0 ; j < size ; j++ {
			player := players[i*size + j]
			team := teams[j]
			team.players = append(team.players, player)
			team.rating = (team.rating * (float32(len(team.players)) - 1) + player.Rating ) / float32(len(team.players))
			teams[j] = team
		}

		sort.Slice(teams[:], func(i, j int) bool {
			return teams[i].rating < teams[j].rating
		})
	}

	return teams
}