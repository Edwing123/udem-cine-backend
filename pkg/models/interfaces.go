package models

type Auth interface {
	Authenticate(Credentials) (int, error)
}

type Movies interface {
	Get(id int) (Movie, error)
	List() ([]Movie, error)
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

type Rooms interface {
	Get(id int) (Room, error)
	List() ([]Room, error)
	ListSeats(id int) ([]Seat, error)
	Create(room NewRoom) error
	Edit(id int, room UpdateRoom) error
	Delete(id int) error
}

type Models struct {
	Auth
	Movies
	Users
	Rooms
}
