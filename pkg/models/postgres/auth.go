package postgres

import (
	"errors"

	"github.com/Edwing123/udem-cine/pkg/hashing"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthController struct {
	conn *pgxpool.Pool
}

func (c *AuthController) Authenticate(credentials models.Credentials) (int, error) {
	var hashedPassword string
	var id int

	row := c.conn.QueryRow(globalCtx, selectIdPassword, credentials.UserName)

	err := row.Scan(&id, &hashedPassword)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, models.ErrAuth
	}

	isValidPassword := hashing.VerifyPassword(credentials.Password, hashedPassword)
	if !isValidPassword {
		return 0, models.ErrAuth
	}

	return id, nil
}
