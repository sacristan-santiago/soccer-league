package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
    "math/rand"
    "time"
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
    var team string
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

func getHandler(selectionFunc func(id int) interface{}) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        if req.Method ==  http.MethodGet {
            idString := req.PathValue("id")
            id, err := strconv.Atoi(idString)
            if err != nil {
                http.Error(w, "Invalid ID", http.StatusBadRequest)
                return
            }

            bodyObj := selectionFunc(id)

            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            err = json.NewEncoder(w).Encode(bodyObj)
            if err != nil {
                http.Error(w, "Cannot parse response", http.StatusInternalServerError)
                return
            }
        }
    }
}

func postHandler(addFunc func(body []byte) int) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        if req.Method ==  http.MethodPost {
            var rawBody json.RawMessage
            err := json.NewDecoder(req.Body).Decode(&rawBody)
            if err != nil {
                fmt.Println(err)
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
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "{\n\"id\": %d\n}", id)
        }
    }
}

func main() {
	db = openDB()
	defer db.Close()

    mux := http.NewServeMux()

    mux.HandleFunc("/players/{id}", getHandler(playerSelectionFunc))

    mux.HandleFunc("/teams/{id}", getHandler(teamSelectionFunc))

    mux.HandleFunc("/match/{id}", getHandler(matchSelectionFunc))

    mux.HandleFunc("/players", postHandler(playerAddFunc))

    mux.HandleFunc("/teams", postHandler(teamAddFunc))

    mux.HandleFunc("/match", postHandler(matchAddFunc))

	fmt.Println("Open your web browser and visit http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}