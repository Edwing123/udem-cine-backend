package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersController struct {
	conn *pgxpool.Pool
}

func (mc *UsersController) Details(id int) (models.User, error) {
	return models.User{}, nil
}

func (mc *UsersController) List() []models.User {
	return []models.User{}
}

func (mc *UsersController) Create(user models.NewUser) error {
	return nil
}

func (mc *UsersController) Edit(user models.User) error {
	return nil
}

func (mc *UsersController) Delete(id int) error {
	return nil
}
