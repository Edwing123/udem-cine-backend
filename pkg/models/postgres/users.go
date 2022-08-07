package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/hashing"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersController struct {
	conn *pgxpool.Pool
}

func (c *UsersController) Scan(row pgx.CollectableRow) (models.User, error) {
	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Role, &user.Password)
	return user, err
}

func (c *UsersController) CheckExecError(err error) error {
	if err != nil {
		if isUniqueViolation(err) {
			return codes.UserNameExists
		}

		return serverError(err)
	}

	return nil
}

func (c *UsersController) Get(id int) (models.User, error) {
	result, err := c.conn.Query(globalCtx, selectUser, id)
	if err != nil {
		return models.User{}, serverError(err)
	}

	user, err := pgx.CollectOneRow(result, c.Scan)
	if isPgxNoRows(err) {
		return user, codes.NoRecords
	}

	return user, nil
}

func (c *UsersController) List() ([]models.User, error) {
	return queryRows(
		c.conn,
		selectAllUsers,
		func(row pgx.CollectableRow) (models.User, error) {
			user, err := c.Scan(row)
			// Don't include password.
			user.Password = ""
			return user, err
		},
	)
}

func (c *UsersController) Create(user models.NewUser) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertUser,
		user.Name,
		user.Role,
		hashing.HashPassword(user.Password),
	)

	return c.CheckExecError(err)
}

func (c *UsersController) Edit(id int, user models.UpdateUser) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateUser,
		id,
		user.Name,
		user.Role,
	)

	return c.CheckExecError(err)
}

func (c *UsersController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteUser,
		id,
	)

	if err != nil {
		return serverError(err)
	}

	return nil
}
