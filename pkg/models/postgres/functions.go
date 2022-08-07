package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FunctionsController struct {
	conn *pgxpool.Pool
}

func (c *FunctionsController) CheckExecError(err error) error {
	if err != nil {
		if isUniqueViolation(err) {
			return codes.FunctionRoomScheduleConflict
		}

		return serverError(err)
	}

	return nil
}

func (c *FunctionsController) Get(id int) (models.Function, error) {
	var function models.Function

	result := c.conn.QueryRow(globalCtx, selectFunction, id)
	err := result.Scan(
		&function.Id,
		&function.Price,
		&function.CreatedAt,
		&function.MovieId,
		&function.Room,
		&function.ScheduleId,
	)

	if isPgxNoRows(err) {
		return function, codes.NoRecords
	}

	return function, nil
}

func (c *FunctionsController) List() ([]models.FunctionDetails, error) {
	return queryRows(c.conn, selectFunctionDetails, func(row pgx.CollectableRow) (models.FunctionDetails, error) {
		var function models.FunctionDetails

		err := row.Scan(
			&function.Id,
			&function.Price,
			&function.CreatedAt,
			&function.Movie,
			&function.Room,
			&function.Schedule,
		)

		return function, err
	})
}

func (c *FunctionsController) Create(function models.NewFunction) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertFunction,
		function.Price,
		function.MovieId,
		function.Room,
		function.ScheduleId,
	)

	return c.CheckExecError(err)
}

func (c *FunctionsController) Edit(id int, function models.UpdateFunction) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateFunction,
		id,
		function.Price,
		function.MovieId,
		function.Room,
		function.ScheduleId,
	)

	return c.CheckExecError(err)
}

func (c *FunctionsController) Archive(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteFunction,
		id,
	)

	if err != nil {
		return serverError(err)
	}

	return nil
}
