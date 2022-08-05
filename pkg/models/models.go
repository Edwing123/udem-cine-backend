package models

import "time"

// Movie related structs.
type Movie struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Classification string `json:"classification"`
	Genre          string `json:"genre"`
	Duration       int    `json:"duration"`
	ReleaseDate    string `json:"releaseDate"`
}

type NewMovie struct {
	Title          string `json:"title"`
	Classification string `json:"classification"`
	Genre          string `json:"genre"`
	Duration       int    `json:"duration"`
	ReleaseDate    string `json:"releaseDate"`
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
	Room   int `json:"room"`
	Number int `json:"number"`
}

// Schedule related structs.
type Schedule struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}

type NewSchedule struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type UpdateSchedule NewSchedule

// Functions related structs.
type Function struct {
	Id         int       `json:"id"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"createdAt"`
	MovieId    int       `json:"movieId"`
	Room       int       `json:"room"`
	ScheduleId int       `json:"scheduleId"`
}

type FunctionDetails struct {
	Id        int       `json:"id"`
	Movie     string    `json:"movie"`
	Schedule  string    `json:"schedule"`
	Room      int       `json:"room"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewFunction struct {
	MovieId    int
	ScheduleId int
	Room       int
	Price      int
}

type UpdateFunction struct {
	MovieId    int
	ScheduleId int
	Room       int
	Price      int
}
