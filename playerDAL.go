package main

import("fmt")

func selectPlayer(id int) Player {
	var player Player 
	row := db.QueryRow(`SELECT * FROM players WHERE id = ?`, id)
	err := row.Scan(&player.Id, &player.FirstName, &player.LastName, &player.Rating)
	handleError(err)
	return player
}

func addPlayer(newPlayer PlayerDTO) int {
    result, err := db.Exec(`INSERT INTO players (first_name, last_name, rating) VALUES (?, ?, ?)`, newPlayer.firstName, newPlayer.lastName, newPlayer.rating)
    handleError(err)

    lastInsertedID, err := result.LastInsertId()
    handleError(err)

    return int(lastInsertedID)
}


func deletePlayer(id int) {
	_, err := db.Exec("DELETE FROM players WHERE id=?", id)
	handleError(err)

	fmt.Println("Succesfully deleted player!")
}

