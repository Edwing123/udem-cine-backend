package main

import (
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/Edwing123/udem-cine/pkg/models/postgres"
)

/*
	Database setup functions and utils.
*/

func NewDatabase(dsn string) models.Models {
	db, err := postgres.New(dsn)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	return db
}
