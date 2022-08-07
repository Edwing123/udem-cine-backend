package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoomsController struct {
	conn *pgxpool.Pool
}

func (c *RoomsController) Scan(row pgx.CollectableRow) (models.Room, error) {
	var room models.Room

	err := row.Scan(
		&room.Number,
		&room.Seats,
	)

	return room, err
}

func (c *RoomsController) ScanSeat(row pgx.CollectableRow) (models.Seat, error) {
	var seat models.Seat
	err := row.Scan(&seat.Number, &seat.Room)
	return seat, err
}

func (c *RoomsController) Get(id int) (models.Room, error) {
	result, err := c.conn.Query(globalCtx, selectRoom, id)
	if err != nil {
		return models.Room{}, serverError(err)
	}

	room, err := pgx.CollectOneRow(result, c.Scan)

	if isPgxNoRows(err) {
		return room, codes.NoRecords
	}

	return room, nil
}

func (c *RoomsController) List() ([]models.Room, error) {
	return queryRows(c.conn, selectAllRooms, c.Scan)
}

func (c *RoomsController) ListSeats(id int) ([]models.Seat, error) {
	return queryRows(c.conn, selectAllSeats, c.ScanSeat)
}

func (c *RoomsController) Create(room models.NewRoom) error {
	tx, err := c.conn.Begin(globalCtx)
	if err != nil {
		return serverError(err)
	}
	defer tx.Rollback(globalCtx)

	_, err = tx.Exec(
		globalCtx,
		insertRoom,
		room.Number,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return codes.RoomExists
		}

		return serverError(err)
	}

	for i := 0; i < room.Seats; i++ {
		_, err := tx.Exec(globalCtx, insertSeat, i+1, room.Number)
		if err != nil {
			return serverError(err)
		}
	}

	err = tx.Commit(globalCtx)
	if err != nil {
		return serverError(err)
	}

	return nil
}

func (c *RoomsController) Edit(id int, room models.UpdateRoom) error {
	tx, err := c.conn.Begin(globalCtx)
	if err != nil {
		return serverError(err)
	}
	defer tx.Rollback(globalCtx)

	err = c.Delete(id)

	if err != nil {
		return err
	}

	err = c.Create(models.NewRoom{
		Number: room.Number,
		Seats:  room.Seats,
	})

	if err != nil {
		return err
	}

	err = tx.Commit(globalCtx)
	if err != nil {
		return serverError(err)
	}

	return nil
}

func (c *RoomsController) Delete(id int) error {
	_, err := c.conn.Exec(
		globalCtx,
		deleteRoom,
		id,
	)

	if err != nil {
		if isFKVilation(err) {
			return codes.FunctionDependsOnRoom
		}

		return serverError(err)
	}

	return nil
}
