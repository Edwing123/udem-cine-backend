package models

import "time"

type Movie struct {
	Id            int
	Title         string
	Clasification string
	Genre         string
	Duration      time.Duration
	ReleaseDate   time.Time
}

type NewMovie struct {
	Title         string
	Clasification string
	Genre         string
	Duration      time.Duration
	ReleaseDate   time.Time
}

type User struct {
	Id       int
	Name     string
	Role     string
	Password string
}

type NewUser struct {
	Name     string
	Role     string
	Password string
}

type Credentials struct {
	Name     string
	Password string
}
