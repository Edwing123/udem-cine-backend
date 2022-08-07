package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/hashing"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthController struct {
	conn *pgxpool.Pool
}

type idAnsPassword struct {
	Id       int
	Password string
}

func (c *AuthController) ScanIdAndPassword(row pgx.CollectableRow) (idAnsPassword, error) {
	var v idAnsPassword
	err := row.Scan(&v.Id, &v.Password)
	return v, err
}

// Authenticates user by checking its credentials from
// the ones of the database.
func (c *AuthController) Authenticate(credentials models.Credentials) (int, error) {
	result, err := c.conn.Query(globalCtx, selectUserIdPassword, credentials.UserName)
	if err != nil {
		return 0, serverError(err)
	}

	record, err := pgx.CollectOneRow(result, c.ScanIdAndPassword)

	if isPgxNoRows(err) {
		return 0, codes.AuthFailed
	}

	isValidPassword := hashing.VerifyPassword(credentials.Password, record.Password)
	if !isValidPassword {
		return 0, codes.AuthFailed
	}

	return record.Id, nil
}
