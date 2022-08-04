package postgres

import (
	"errors"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RoomsController struct {
	conn *pgxpool.Pool
}

func (c *RoomsController) Get(id int) (models.Room, error) {
	var room models.Room

	result := c.conn.QueryRow(globalCtx, selectRoom, id)
	err := result.Scan(
		&room.Number,
		&room.Seats,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return room, models.ErroNoRows
	}

	return room, nil
}

func (c *RoomsController) List() ([]models.Room, error) {
	rooms := make([]models.Room, 0)

	result, err := c.conn.Query(globalCtx, selectAllRooms)
	if err != nil {
		return nil, serverError(err)
	}

	for result.Next() {
		var room models.Room

		err := result.Scan(
			&room.Number,
			&room.Seats,
		)

		if err != nil {
			return nil, serverError(err)
		}

		rooms = append(rooms, room)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return rooms, nil
}

func (c *RoomsController) ListSeats(id int) ([]models.Seat, error) {
	seats := make([]models.Seat, 0)

	result, err := c.conn.Query(globalCtx, selectAllSeats, id)
	if err != nil {
		return nil, serverError(err)
	}

	for result.Next() {
		var seat models.Seat

		err := result.Scan(
			&seat.Number,
			&seat.Room,
		)

		if err != nil {
			return nil, serverError(err)
		}

		seats = append(seats, seat)
	}

	if err := result.Err(); err != nil {
		return nil, serverError(err)
	}

	return seats, nil
}

func (c *RoomsController) Create(room models.NewRoom) error {
	tx, err := c.conn.Begin(globalCtx)

	if err != nil {
		return serverError(err)
	}

	// Defer call to rollback to cancel the transaction if
	// something goes wrong.
	defer tx.Rollback(globalCtx)

	_, err = tx.Exec(
		globalCtx,
		insertRoom,
		room.Number,
	)

	if err != nil {
		// Is room already taken?
		if getPgxErroCode(err) == pgerrcode.UniqueViolation {
			return models.ErrRoomTaken
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

	// Defer call to rollback to cancel the transaction if
	// something goes wrong.
	defer tx.Rollback(globalCtx)

	err = c.Delete(id)

	if err != nil {
		return serverError(err)
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
		return serverError(err)
	}

	return nil
}
