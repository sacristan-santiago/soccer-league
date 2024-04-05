package main

func selectMatch(id int) Match {
	var match Match 
	row := db.QueryRow(`SELECT * FROM matches WHERE id = ?`, id)
	err := row.Scan(&match.Id, &match.TeamA, &match.TeamB, &match.ScoreA, &match.ScoreB)
	handleError(err)
	return match
}

func addMatch(newMatch MatchDTO) int {
	result, err := db.Exec(`INSERT INTO matches (teamA, teamB, scoreA, scoreB) VALUES (?, ?, ?, ?)`, newMatch.TeamA, newMatch.TeamB, newMatch.ScoreA ,newMatch.ScoreB)
	handleError(err)
	
	lastInsertedID, err := result.LastInsertId()
	handleError(err)

	return int(lastInsertedID)
}

func deleteMatch(id int) {
	_, err := db.Exec("DELETE FROM matches WHERE id=?", id)
	handleError(err)
}