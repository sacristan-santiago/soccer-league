package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // Declare db globally

func openDB() *sql.DB  {
	var err error
	db, err = sql.Open("sqlite3", "./soccer-league.db")
	handleError(err)

	_, err = db.Exec(`

					PRAGMA foreign_keys = ON;
					
					CREATE TABLE IF NOT EXISTS players (
						id INTEGER PRIMARY KEY,
						first_name TEXT,
						last_name TEXT,
						rating REAL
					);
					CREATE TABLE IF NOT EXISTS matches (
						id INTEGER PRIMARY KEY,
						teamA INTEGER,
						teamB INTEGER,
						scoreA INTEGER,
						scoreB INTEGER,
						FOREIGN KEY (teamA) REFERENCES teams(id),
						FOREIGN KEY (teamB) REFERENCES teams(id)
					);
					
					CREATE TABLE IF NOT EXISTS teams (
						id INTEGER PRIMARY KEY,
						name TEXT
					);

					CREATE TABLE IF NOT EXISTS team_player(
						id INTEGER PRIMARY KEY,
						team_id INTEGER,
						player_id INTEGER,
						FOREIGN KEY (team_id) REFERENCES teams(id),
						FOREIGN KEY (player_id) REFERENCES players(id)
					)
					`)
	handleError(err)

	return db
}