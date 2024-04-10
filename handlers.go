package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

func postDeleteTeamPlayerHandler() http.HandlerFunc {
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

func getAllEntitiesHandler(selectData func() interface{}) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        if req.Method == http.MethodGet {
            data := selectData()

            err := json.NewEncoder(w).Encode(data)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}

func getPlayersWrapper() interface{}{
    return selectPlayers()
}
func getTeamsWrapper() interface{}{
    return selectTeams()
}
func getMatchesWrapper() interface{}{
    return selectMatches()
}

func getPlayerTeams() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")

		playerIDString := req.PathValue("id")
        playerID, err := strconv.Atoi(playerIDString)
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }
        
        if req.Method ==  http.MethodGet {
            teams := selectPlayerTeams(playerID)

            err := json.NewEncoder(w).Encode(teams)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}


func getDeleteUpdateHandler(selectionFunc func(id int) interface{}, deleteFunc func(id int), updateFunc func(id int, body []byte) interface{}) http.HandlerFunc {
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
        } else if req.Method == http.MethodPut {
            var rawBody json.RawMessage
            err := json.NewDecoder(req.Body).Decode(&rawBody)
            if err != nil {
                http.Error(w, "Body cannot be parsed", http.StatusBadRequest)
                return
            }
            
            updatedEntity := updateFunc(id, rawBody)

            err = json.NewEncoder(w).Encode(updatedEntity)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}

func updateWrapperMock(id int, body []byte) interface{} {return nil}

func updateMatchWrapper(id int, body []byte) interface{} {
    var match MatchDTO
    err := json.Unmarshal(body, &match)
    if err != nil {
        panic(err)
    }

    return updateMatchScore(id, match)
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
        deletePlayerfromTeam(id, team.Id)
    }

    deletePlayer(id)
}