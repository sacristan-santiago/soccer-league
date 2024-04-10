package main

import (
	"fmt"
	"net/http"
)

func main() {
	db = openDB()
	defer db.Close()

    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir("./static")))
    mux.HandleFunc("/players/{id}", getDeleteUpdateHandler(playerSelectionFunc, playerDeleteFunc, updateWrapperMock))
    mux.HandleFunc("/players", postHandler(playerAddFunc))
    mux.HandleFunc("/players/all", getAllEntitiesHandler(getPlayersWrapper))
    mux.HandleFunc("/player/teams/{id}", getPlayerTeams())
    mux.HandleFunc("/teams/{id}", getDeleteUpdateHandler(teamSelectionFunc, deleteTeam, updateWrapperMock))
    mux.HandleFunc("/teams/player/{teamID}/{playerID}", postDeleteTeamPlayerHandler())
    mux.HandleFunc("/teams/all", getAllEntitiesHandler(getTeamsWrapper))
    mux.HandleFunc("/teams", postHandler(teamAddFunc))
    mux.HandleFunc("/teams/player", postHandler(teamPlayerAddFunc))
    mux.HandleFunc("/match", postHandler(matchAddFunc))
    mux.HandleFunc("/match/{id}", getDeleteUpdateHandler(matchSelectionFunc, deleteMatch, updateMatchWrapper))
    mux.HandleFunc("/match/all", getAllEntitiesHandler(getMatchesWrapper))

    fmt.Println("Open your web browser and visit http://localhost:3000")
    http.ListenAndServe(":3000", mux)
}