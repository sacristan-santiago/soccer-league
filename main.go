package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

func main() {
	db = openDB()
	defer db.Close()

    mux := http.NewServeMux()

    mux.HandleFunc("/players/{id}", func(w http.ResponseWriter, req *http.Request) {
        idString := req.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

		player := selectPlayer(id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

        err = json.NewEncoder(w).Encode(player)
        if err != nil {
            http.Error(w, "Cannot parse response", http.StatusInternalServerError)
            return
        }
    })

	fmt.Println("Open your web browser and visit http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}