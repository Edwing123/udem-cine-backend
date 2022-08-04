package models

import "time"

// Movie related structs.
type Movie struct {
	Id             int           `json:"id"`
	Title          string        `json:"title"`
	Classification string        `json:"classification"`
	Genre          string        `json:"genre"`
	Duration       time.Duration `json:"duration"`
	ReleaseDate    time.Time     `json:"releaseDate"`
}

type NewMovie struct {
	Title          string        `json:"title"`
	Classification string        `json:"classification"`
	Genre          string        `json:"genre"`
	Duration       time.Duration `json:"duration"`
	ReleaseDate    time.Time     `json:"releaseDate"`
}

type UpdateMovie NewMovie

// User related structs.
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

type NewUser struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// Credentials struct.
type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Room related structs.
type Room struct {
	Number int `json:"number"`
	Seats  int `json:"seats"`
}

type NewRoom Room

type UpdateRoom Room

type Seat struct {
	Room   int
	Number int
}

// Schedule related structs.
type Schedule struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type NewSchedule struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type UpdateSchedule NewSchedule

// Functions related structs.
type Function struct {
	Id         int
	Price      float64
	CreatedAt  time.Time
	MovieId    int
	Room       int
	ScheduleId int
}

type FunctionDetails struct {
	Id        int
	Movie     string
	Schedule  string
	Room      int
	Price     float64
	createdAt string
}

type NewFunction struct {
	movieId    int
	ScheduleId int
	Room       int
	Price      float64
}

type UpdateFunction struct {
	movieId    int
	ScheduleId int
	Room       int
	Price      float64
}
