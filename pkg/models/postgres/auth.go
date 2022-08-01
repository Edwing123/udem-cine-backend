package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthController struct {
	conn *pgxpool.Pool
}

func (ac *AuthController) Login(models.Credentials) error {
	return nil
}
