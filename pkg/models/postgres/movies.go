package postgres

import (
	"time"

	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MoviesController struct {
	conn *pgxpool.Pool
}

func (c *MoviesController) Scan(row pgx.CollectableRow) (models.Movie, error) {
	var movie models.Movie
	var date time.Time

	err := row.Scan(
		&movie.Id,
		&movie.Title,
		&movie.Classification,
		&movie.Genre,
		&movie.Duration,
		&date,
	)

	movie.ReleaseDate = date.Format("2006-01-02")

	return movie, err
}

func (c *MoviesController) CheckExecError(err error) error {
	if err != nil {
		if isUniqueViolation(err) {
			return codes.MovieTitleExists
		}

		return serverError(err)
	}

	return nil
}

func (c *MoviesController) ParseDate(date string) time.Time {
	d, _ := time.Parse("2006-01-02", date)
	return d
}

func (c *MoviesController) Get(id int) (models.Movie, error) {
	result, err := c.conn.Query(globalCtx, selectMovie, id)
	if err != nil {
		return models.Movie{}, err
	}

	movie, err := pgx.CollectOneRow(result, c.Scan)

	if isPgxNoRows(err) {
		return movie, codes.NoRecords
	}

	return movie, nil
}

func (c *MoviesController) List() ([]models.Movie, error) {
	return queryRows(
		c.conn,
		selectAllMovies,
		c.Scan,
	)
}

func (c *MoviesController) Create(movie models.NewMovie) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertMovie,
		movie.Title,
		movie.Classification,
		movie.Genre,
		movie.Duration,
		movie.ReleaseDate,
	)

	return c.CheckExecError(err)
}

func (c *MoviesController) Edit(id int, movie models.UpdateMovie) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateMovie,
		id,
		movie.Title,
		movie.Classification,
		movie.Genre,
		movie.Duration,
		movie.ReleaseDate,
	)

	return c.CheckExecError(err)
}

func (c *MoviesController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteMovie,
		id,
	)

	if err != nil {
		if isFKVilation(err) {
			return codes.FunctionDependsOnMovie
		}

		return serverError(err)
	}

	return nil
}
