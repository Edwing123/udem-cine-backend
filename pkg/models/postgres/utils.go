package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgconn"
)

var (
	globalCtx = context.Background()
)

func serverError(err error) error {
	return fmt.Errorf("%w: %s", models.ErrServer, err)
}

func getPgxErroCode(err error) string {
	var pgErr *pgconn.PgError
	errors.As(err, &pgErr)
	return pgErr.Code
}
