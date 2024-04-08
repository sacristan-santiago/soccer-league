package main

import("fmt")

func selectPlayers() []Player {
	var players []Player
	rows, err := db.Query(
		`SELECT 
			id,
			first_name,
			last_name,
			rating
		FROM players
		ORDER BY id ASC
		LIMIT 20
		`)
	handleError(err)
	defer rows.Close()

	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.Id, &player.FirstName, &player.LastName, &player.Rating); err != nil {
			handleError(err)
		}
		players = append(players, player)
	}
    if err = rows.Err(); err != nil {
        handleError(err)
    }

	return players
}


func selectPlayer(id int) Player {
	var player Player 
	row := db.QueryRow(`SELECT * FROM players WHERE id = ?`, id)
	err := row.Scan(&player.Id, &player.FirstName, &player.LastName, &player.Rating)
	handleError(err)
	return player
}

func addPlayer(newPlayer PlayerDTO) int {
    result, err := db.Exec(`INSERT INTO players (first_name, last_name, rating) VALUES (?, ?, ?)`, newPlayer.FirstName, newPlayer.LastName, newPlayer.Rating)
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

