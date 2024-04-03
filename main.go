package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

type Player struct {
	id int
	firstName string
	lastName string
	rating float32
}

type PlayerDTO struct {
	firstName string
	lastName string
	rating float32
}

type Team struct {
	id int 
	name string
	players []Player
	rating float32
}

type Match struct {
	id int
	teamA int
	teamB int
	scoreA int
	scoreB int
}

type MatchDTO struct {
	teamA int
	teamB int
	scoreA int
	scoreB int
}


func sortTeams(players []Player, size int) []Team {
	if len(players) % size != 0 {
		panic("Team size results in unassigned players.")
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].rating > players[j].rating
	})
	
	var teams = make([]Team, size)

	for i := 0; i < len(players) / size ; i++ {
		for j := 0 ; j < size ; j++ {
			player := players[i*size + j]
			team := teams[j]
			team.players = append(team.players, player)
			team.rating = (team.rating * (float32(len(team.players)) - 1) + player.rating ) / float32(len(team.players))
			teams[j] = team
		}

		sort.Slice(teams[:], func(i, j int) bool {
			return teams[i].rating < teams[j].rating
		})
	}

	return teams
}

func main() {
	db = openDB()
	defer db.Close()

	var teams []Team
	for i := 0; i < 2; i++ {
		teamId := addTeam(fmt.Sprintf("Team %d", i))

		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			playerId := addPlayer(PlayerDTO{"Player", fmt.Sprintf("%d", (j + 1) + (i * 5)), float32(rand.Intn(10) + 1)})
			addPlayertoTeam(playerId, teamId)
		}

		team := selectTeam(teamId)
		teams = append(teams, team)
		fmt.Println(team)
	}

	matchId := addMatch(MatchDTO{teams[0].id, teams[1].id, 2, 1})
	match := selectMatch(matchId)
	fmt.Println(match)

	//delete all
	for i := 0; i < 2; i++ {
		deleteTeam(teams[i].id)
		fmt.Println("Deleted team:", teams[i].id)

		for j := 0; j < 5; j++ {
			deletePlayer(j + (5 * i) + 1)
			fmt.Println("Deleted player:", j + (5 * i) + 1)
		}
	}

	deleteMatch(matchId)
	fmt.Println("Deleted match:", matchId)
}