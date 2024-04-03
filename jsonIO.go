package main

import (
	"encoding/json"
	"os"
	"io"
)

func openFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	handleError(err)
	return file
}

func decode(file *os.File, players *map[string]Player) {
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&players)
	if err != nil && err != io.EOF {
		panic(err)
	}
}

func encode (file *os.File, players map[string]Player) {
	file.Truncate(0)
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	err := encoder.Encode(players)
	handleError(err)
}

// func addPlayer(newPlayer Player) string {
// 	players := make(map[string]Player)

// 	file := openFile("players.json")
// 	defer file.Close()

// 	decode(file, &players)

// 	id := generateUUID()
// 	players[id] = newPlayer

// 	encode(file, players)

// 	fmt.Println("New player added successfully.")

// 	return id
// }

// func deletePlayer(id string) {
// 	players := make(map[string]Player)

// 	file := openFile("players.json")
// 	defer file.Close()

// 	decode(file, &players)

// 	delete(players, id)

// 	encode(file, players)

// 	fmt.Println("New player deleted successfully.")
// }

// func selectPlayer(id string) Player {
// 	players := make(map[string]Player)

// 	file := openFile("players.json")
// 	defer file.Close()

// 	decode(file, &players)

// 	return players[id]
// }

// func selectPlayers() []Player {
// 	players := make(map[string]Player)
	 

// 	file := openFile("players.json")
// 	defer file.Close()

// 	decode(file, &players)

// 	var playersSlice []Player
// 	for _, v := range players {
// 		playersSlice = append(playersSlice, v)
// 	}

// 	return playersSlice
// }
