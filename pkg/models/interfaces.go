package models

type Auth interface {
	Authenticate(Credentials) (int, error)
}

type Movies interface {
	List() []Movie
	Create(movie NewMovie) error
	Edit(id int, updateMovie UpdateMovie) error
	Delete(id int) error
}

type Users interface {
	Get(id int) (User, error)
	List() ([]User, error)
	Create(user NewUser) error
	Edit(id int, user UpdateUser) error
	Delete(id int) error
}

type Models struct {
	Auth
	Movies
	Users
}
