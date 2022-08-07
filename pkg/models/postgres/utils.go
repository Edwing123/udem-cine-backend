package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	globalCtx = context.Background()
)

type RowScanner interface {
	Scan(...any) error
}

func serverError(err error) error {
	return fmt.Errorf("%w: %s", codes.ServerFailed, err)
}

func getPgxErroCode(err error) string {
	var pgErr *pgconn.PgError
	errors.As(err, &pgErr)
	return pgErr.Code
}

func isUniqueViolation(err error) bool {
	return getPgxErroCode(err) == pgerrcode.UniqueViolation
}

func isFKVilation(err error) bool {
	return getPgxErroCode(err) == pgerrcode.ForeignKeyViolation
}

func isPgxNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func queryRows[T any](conn *pgxpool.Pool, query string, scanner pgx.RowToFunc[T]) ([]T, error) {
	result, err := conn.Query(globalCtx, query)
	if err != nil {
		return []T{}, serverError(err)
	}

	values, err := pgx.CollectRows(result, scanner)
	if err := result.Err(); err != nil {
		return values, serverError(err)
	}

	return values, nil
}
