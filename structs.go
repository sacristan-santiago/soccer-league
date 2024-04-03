package main

type Player struct {
	Id int
	FirstName string
	LastName string
	Rating float32
}

type PlayerDTO struct {
	firstName string
	lastName string
	rating float32
}

type Team struct {
	id int 
	name string
	players []Player
	rating float32
}

type Match struct {
	id int
	teamA int
	teamB int
	scoreA int
	scoreB int
}

type MatchDTO struct {
	teamA int
	teamB int
	scoreA int
	scoreB int
}