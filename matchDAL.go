package main

func selectMatches() []Match {
	var matches []Match
	rows, err := db.Query(
		`SELECT
			id,
			teamA,
			teamB,
			scoreA,
			scoreB
		FROM matches
		ORDER BY id ASC
		LIMIT 20
		`)
	handleError(err)
	defer rows.Close()

	for rows.Next() {
		var match Match
		if err := rows.Scan(&match.Id, &match.TeamA, &match.TeamB, &match.ScoreA, &match.ScoreB); err != nil {
			handleError(err)
		}
		matches = append(matches, match)
	}
    if err = rows.Err(); err != nil {
        handleError(err)
    }

	return matches
}

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

func updateMatchScore(id int, match MatchDTO) Match{
	_,err := db.Exec("UPDATE matches SET scoreA = ?, scoreB = ? WHERE id = ?", match.ScoreA, match.ScoreB, id)
	handleError(err)
	
	return selectMatch(id)
}