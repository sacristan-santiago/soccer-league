package main

func selectTeam(id int) Team {
	var team Team
	rows, err := db.Query(
		`SELECT 
			teams.id, 
			teams.name,
			players.id,
			players.first_name,
			players.last_name,
			players.rating			
		FROM teams 
		JOIN team_player ON teams.id = team_player.team_id
		JOIN players ON players.id = team_player.player_id
		WHERE teams.id = ?
		`, id)
	handleError(err)
	defer rows.Close()

	for rows.Next() {
		var player Player
		if err := rows.Scan(&team.id, &team.name, &player.Id, &player.FirstName, &player.LastName, &player.Rating); err != nil {
			handleError(err)
		}
		team.players = append(team.players, player)
	}
    if err = rows.Err(); err != nil {
        handleError(err)
    }

	row := db.QueryRow(
		`SELECT AVG(players.rating)
		FROM players 
		JOIN team_player ON players.id = team_player.player_id 
		WHERE team_player.team_id = ?
		`, id)
	err = row.Scan(&team.rating)
	handleError(err)

	return team
}

func addTeam(name string) int {
    result, err := db.Exec(`INSERT INTO teams (name) VALUES (?)`, name)
    handleError(err)

    lastInsertedID, err := result.LastInsertId()
    handleError(err)

    return int(lastInsertedID)
}

func deleteTeam(id int) {
	_, err := db.Exec(`
		DELETE FROM team_player where team_id = ?;
		DELETE FROM teams WHERE id=?;
	`, id, id)
	handleError(err)

}

func addPlayertoTeam(playerId int, teamId int) int {
	result, err := db.Exec(`INSERT INTO team_player (player_id, team_id) VALUES (?, ?)`, playerId, teamId)
    handleError(err)

    lastInsertedID, err := result.LastInsertId()
    handleError(err)

    return int(lastInsertedID)	
}

func deletePlayerfromTeam(playerId int, teamId int) {
	_, err := db.Exec(`DELETE FROM team_player WHERE player_id = ? AND team_id = ? `, playerId, teamId)
	handleError(err)
}