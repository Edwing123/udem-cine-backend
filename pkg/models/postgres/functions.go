package postgres

import (
	"errors"
	"fmt"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FunctionsController struct {
	conn *pgxpool.Pool
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

	fmt.Println(result, err)
	fmt.Println(function)

	if errors.Is(err, pgx.ErrNoRows) {
		return function, models.ErroNoRows
	}

	return function, nil
}

func (c *FunctionsController) List() ([]models.FunctionDetails, error) {
	functions := make([]models.FunctionDetails, 0)

	result, err := c.conn.Query(globalCtx, selectFunctionDetails)

	if err != nil {
		return nil, serverError(err)
	}

	for result.Next() {
		var function models.FunctionDetails

		err := result.Scan(
			&function.Id,
			&function.Price,
			&function.CreatedAt,
			&function.Movie,
			&function.Room,
			&function.Schedule,
		)

		fmt.Println(err)

		if err != nil {
			return nil, serverError(err)
		}

		functions = append(functions, function)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return functions, nil
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

	if err != nil {
		// Check for unique contraints violation...
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrFunctionFuckedUp
		}

		return serverError(err)
	}

	return nil
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

	if err != nil {
		// Check for unique contraints violation...
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrFunctionFuckedUp
		}

		return serverError(err)
	}

	return nil
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
