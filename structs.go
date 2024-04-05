package main

type Player struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Rating float32 `json:"rating"`
}

type PlayerDTO struct {
	FirstName string
	LastName  string
	Rating    float32
}

type Team struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Players []Player `json:"players"`
	Rating float32 `json:"rating"`
}

type Match struct {
	Id int `json:"id"`
	TeamA int `json:"teamA"`
	TeamB int `json:"teamB"`
	ScoreA int `json:"scoreA"`
	ScoreB int `json:"scoreB"`
}

type MatchDTO struct {
	TeamA int
	TeamB int
	ScoreA int
	ScoreB int
}