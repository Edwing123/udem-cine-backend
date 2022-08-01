package models

type Auth interface {
	Login(Credentials) error
}

type Movies interface {
	List() []Movie
	Create(movie NewMovie) error
	Edit(Movie) error
	Delete(id int) error
}

type Users interface {
	Details(id int) (User, error)
	List() []User
	Create(user NewUser) error
	Edit(user User) error
	Delete(id int) error
}

type Models struct {
	Auth
	Movies
	Users
}
