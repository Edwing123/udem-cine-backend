package postgres

import (
	"errors"

	"github.com/Edwing123/udem-cine/pkg/hashing"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersController struct {
	conn *pgxpool.Pool
}

func (c *UsersController) Get(id int) (models.User, error) {
	var user models.User

	result := c.conn.QueryRow(globalCtx, selectUser, id)
	err := result.Scan(&user.Id, &user.Name, &user.Role, &user.Password)

	if errors.Is(err, pgx.ErrNoRows) {
		return user, models.ErroNoRows
	}

	return user, nil
}

func (c *UsersController) List() ([]models.User, error) {
	users := make([]models.User, 0)

	result, err := c.conn.Query(globalCtx, selectAllUsers)
	if err != nil {
		return nil, serverError(err)
	}

	for result.Next() {
		var user models.User

		err := result.Scan(&user.Id, &user.Name, &user.Role, nil)
		if err != nil {
			return nil, serverError(err)
		}

		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return users, nil
}

func (c *UsersController) Create(user models.NewUser) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertUser,
		user.Name,
		user.Role,
		hashing.HashPassword(user.Password),
	)

	if err != nil {
		// Is name already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrUserNameTaken
		}

		return serverError(err)
	}

	return nil
}

func (c *UsersController) Edit(id int, user models.UpdateUser) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateUser,
		id,
		user.Name,
		user.Role,
	)
	if err != nil {
		// Is name already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrUserNameTaken
		}

		return serverError(err)
	}

	return nil
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
