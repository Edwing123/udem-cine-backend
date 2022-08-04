package postgres

import (
	"errors"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SchedulesController struct {
	conn *pgxpool.Pool
}

func (c *SchedulesController) Get(id int) (models.Schedule, error) {
	var schedule models.Schedule

	result := c.conn.QueryRow(globalCtx, selectSchedule, id)
	err := result.Scan(
		&schedule.Id,
		&schedule.Name,
		&schedule.Time,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return schedule, models.ErroNoRows
	}

	return schedule, nil
}

func (c *SchedulesController) List() ([]models.Schedule, error) {
	var schedules []models.Schedule

	result, err := c.conn.Query(globalCtx, selectAllSchedules)
	if err != nil {
		return nil, serverError(err)
	}

	for result.Next() {
		var schedule models.Schedule

		err := result.Scan(
			&schedule.Id,
			&schedule.Name,
			&schedule.Time,
		)

		if err != nil {
			return nil, serverError(err)
		}

		schedules = append(schedules, schedule)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return schedules, nil
}

func (c *SchedulesController) Create(schedule models.NewSchedule) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertSchedule,
		&schedule.Name,
		&schedule.Time,
	)

	if err != nil {
		// Is schedule time already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrScheduleTaken
		}

		return serverError(err)
	}

	return nil
}

func (c *SchedulesController) Edit(id int, schedule models.UpdateSchedule) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateSchedule,
		id,
		&schedule.Name,
		&schedule.Time,
	)

	if err != nil {
		// Is schedule time already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrScheduleTaken
		}

		return serverError(err)
	}

	return nil
}

func (c *SchedulesController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteSchedule,
		id,
	)

	if err != nil {
		return serverError(err)
	}

	return nil
}
