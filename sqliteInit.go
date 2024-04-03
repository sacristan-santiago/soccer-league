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
						scoreB INTEGER
					);
					
					CREATE TABLE IF NOT EXISTS teams (
						id INTEGER PRIMARY KEY,
						name TEXT
					);

					CREATE TABLE IF NOT EXISTS team_player(
						id INTEGER PRIMARY KEY,
						team_id INTEGER,
						player_id INTEGER
					)
					`)
	handleError(err)

	return db
}