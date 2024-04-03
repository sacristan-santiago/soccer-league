package main

func selectMatch(id int) Match {
	var match Match 
	row := db.QueryRow(`SELECT * FROM matches WHERE id = ?`, id)
	err := row.Scan(&match.id, &match.teamA, &match.teamB, &match.scoreA, &match.scoreB)
	handleError(err)
	return match
}

func addMatch(newMatch MatchDTO) int {
	result, err := db.Exec(`INSERT INTO matches (teamA, teamB, scoreA, scoreB) VALUES (?, ?, ?, ?)`, newMatch.teamA, newMatch.teamB, newMatch.scoreA ,newMatch.scoreB)
	handleError(err)
	
	lastInsertedID, err := result.LastInsertId()
	handleError(err)

	return int(lastInsertedID)
}

func deleteMatch(id int) {
	_, err := db.Exec("DELETE FROM matches WHERE id=?", id)
	handleError(err)
}