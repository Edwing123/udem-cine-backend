package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SchedulesController struct {
	conn *pgxpool.Pool
}

func (c *SchedulesController) Scan(row pgx.CollectableRow) (models.Schedule, error) {
	var schedule models.Schedule

	err := row.Scan(
		&schedule.Id,
		&schedule.Name,
		&schedule.Time,
	)

	return schedule, err
}

func (c *SchedulesController) CheckExecError(err error) error {
	if err != nil {
		if isUniqueViolation(err) {
			return codes.ScheduleExists
		}

		return serverError(err)
	}

	return nil
}

func (c *SchedulesController) Get(id int) (models.Schedule, error) {
	result, err := c.conn.Query(globalCtx, selectSchedule, id)
	if err != nil {
		return models.Schedule{}, serverError(err)
	}

	schedule, err := pgx.CollectOneRow(result, c.Scan)

	if isPgxNoRows(err) {
		return schedule, codes.NoRecords
	}

	return schedule, nil
}

func (c *SchedulesController) List() ([]models.Schedule, error) {
	return queryRows(c.conn, selectAllSchedules, c.Scan)
}

func (c *SchedulesController) Create(schedule models.NewSchedule) error {
	_, err := c.conn.Exec(
		globalCtx,
		insertSchedule,
		&schedule.Name,
		&schedule.Time,
	)

	return c.CheckExecError(err)
}

func (c *SchedulesController) Edit(id int, schedule models.UpdateSchedule) error {
	_, err := c.conn.Exec(
		globalCtx,
		updateSchedule,
		id,
		&schedule.Name,
		&schedule.Time,
	)

	return c.CheckExecError(err)
}

func (c *SchedulesController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteSchedule,
		id,
	)

	if err != nil {
		if isFKVilation(err) {
			return codes.FunctionDependsOnSchedule
		}

		return serverError(err)
	}

	return nil
}
