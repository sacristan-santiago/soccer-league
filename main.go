package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)


func playerSelectionFunc(id int) interface{} {
    return selectPlayer(id)
}

func teamSelectionFunc(id int) interface{} {
    return selectTeam(id)
}

func matchSelectionFunc(id int) interface{} {
    return selectMatch(id)
}

func playerAddFunc(body []byte) int {
    var player PlayerDTO
    err := json.Unmarshal(body, &player)
    if err != nil {
        panic(err)
    }

    return addPlayer(player)
}

func teamAddFunc(body []byte) int {
    var team TeamDTO
    err := json.Unmarshal(body, &team)
    if err != nil {
        panic(err)
    }

    return addTeam(team)
}

func matchAddFunc(body []byte) int {
    var match MatchDTO
    err := json.Unmarshal(body, &match)
    if err != nil {
        panic(err)
    }

    return addMatch(match)
}

func teamPlayerAddFunc(body []byte) int {
    var teamPlayer TeamPlayerDTO
    if err := json.Unmarshal(body, &teamPlayer); err != nil {
        panic(err)
    }

    return addPlayertoTeam(teamPlayer.PlayerId, teamPlayer.TeamId)
}

func playerDeleteFunc(id int) {
    playerTeams := selectPlayerTeams(id)

    for _, team := range playerTeams {
        deletePlayerfromTeam(id, team)
    }

    deletePlayer(id)
}

func deleteTeamPlayerHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        teamIDString := req.PathValue("teamID")
        teamID, err := strconv.Atoi(teamIDString)
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        playerIDString := req.PathValue("playerID")
        playerID, err := strconv.Atoi(playerIDString)
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }
        
        if req.Method == http.MethodPost {
            addPlayertoTeam(playerID, teamID)
        } else if req.Method == http.MethodDelete {
            deletePlayerfromTeam(playerID, teamID)
            fmt.Fprintf(w, "Player with ID: %d succesfully removed from team: %d", playerID, teamID )
        }
    }
}

func getTeamsHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        if req.Method ==  http.MethodGet {
            teams := selectTeams()

            err := json.NewEncoder(w).Encode(teams)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}

func getPlayersHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        if req.Method ==  http.MethodGet {
            players := selectPlayers()

            err := json.NewEncoder(w).Encode(players)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}


func getDeleteHandler(selectionFunc func(id int) interface{}, deleteFunc func(id int)) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        idString := req.PathValue("id")
        id, err := strconv.Atoi(idString)
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }
        
        if req.Method ==  http.MethodGet {
            bodyObj := selectionFunc(id)

            err = json.NewEncoder(w).Encode(bodyObj)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        } else if req.Method == http.MethodDelete {
            deleteFunc(id)
            fmt.Fprintf(w, "Entity with Id: %d Succesfully deleted", id)
        }
    }
}

func postHandler(addFunc func(body []byte) int) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        if req.Method ==  http.MethodPost {
            var rawBody json.RawMessage
            err := json.NewDecoder(req.Body).Decode(&rawBody)
            if err != nil {
                http.Error(w, "Body cannot be parsed", http.StatusBadRequest)
                return
            }

            var id int
            id = addFunc(rawBody)
            if err != nil {
                http.Error(w, "Internal Server error", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            fmt.Fprintf(w, "{\n\"id\": %d\n}", id)
        }
    }
}

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