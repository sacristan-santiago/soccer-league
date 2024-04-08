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
    mux.HandleFunc("/players/{id}", getDeleteHandler(playerSelectionFunc, playerDeleteFunc))
    mux.HandleFunc("/teams/{id}", getDeleteHandler(teamSelectionFunc, deleteTeam))
    mux.HandleFunc("/match/{id}", getDeleteHandler(matchSelectionFunc, deleteMatch))
    mux.HandleFunc("/teams/player/{teamID}/{playerID}", deleteTeamPlayerHandler())
    mux.HandleFunc("/teams/all", getTeamsHandler())
    mux.HandleFunc("/players", postHandler(playerAddFunc))
    mux.HandleFunc("/players/all", getPlayersHandler())
    mux.HandleFunc("/teams", postHandler(teamAddFunc))
    mux.HandleFunc("/match", postHandler(matchAddFunc))
    mux.HandleFunc("/teams/player", postHandler(teamPlayerAddFunc))


    fmt.Println("Open your web browser and visit http://localhost:3000")
    http.ListenAndServe(":3000", mux)
}