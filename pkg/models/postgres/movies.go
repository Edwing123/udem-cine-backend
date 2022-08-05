package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type MoviesController struct {
	conn *pgxpool.Pool
}

func (c *MoviesController) Get(id int) (models.Movie, error) {
	var movie models.Movie

	var releaseDate time.Time

	result := c.conn.QueryRow(globalCtx, selectMovie, id)
	err := result.Scan(
		&movie.Id,
		&movie.Title,
		&movie.Classification,
		&movie.Genre,
		&movie.Duration,
		&releaseDate,
	)

	movie.ReleaseDate = releaseDate.Format("2006-01-02")

	if errors.Is(err, pgx.ErrNoRows) {
		return movie, models.ErroNoRows
	}

	return movie, nil
}

func (c *MoviesController) List() ([]models.Movie, error) {
	movies := make([]models.Movie, 0)

	result, err := c.conn.Query(globalCtx, selectAllMovies)

	if err != nil {
		return nil, serverError(err)
	}

	var releaseDate time.Time

	for result.Next() {
		var movie models.Movie

		err := result.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Classification,
			&movie.Genre,
			&movie.Duration,
			&releaseDate,
		)

		movie.ReleaseDate = releaseDate.Format("02 de January del 2006")

		if err != nil {
			return nil, serverError(err)
		}

		movies = append(movies, movie)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return movies, nil
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

	if err != nil {
		// Is title already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrMovieTitleTaken
		}

		return serverError(err)
	}

	return nil
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

	if err != nil {
		// Is title already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrMovieTitleTaken
		}

		return serverError(err)
	}

	return nil
}

func (c *MoviesController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteMovie,
		id,
	)

	fmt.Println(err)

	if err != nil {
		return serverError(err)
	}

	return nil
}
