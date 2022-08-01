package main

import (
	"flag"
	"log"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Models      models.Models
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (api *Api) NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiber",
		AppName:      "UdeMCine API",
	})

	// Define routes.
	app.Post("/auth/login", api.AuthLogin)
	app.Get("/user", api.UserDetaills)

	users := app.Group("/users")
	users.Get("/list", api.UsersList)
	users.Patch("/edit", api.UsersEdit)
	users.Delete("/delete", api.UsersDelete)

	movies := app.Group("/movies")
	movies.Get("/list", api.MoviesList)
	movies.Patch("/edit", api.MoviesEdit)
	movies.Delete("/delete", api.MoviesDelete)

	return app
}

// Command line arguments.
type Args struct {
	Dsn     string
	Address string
}

func GetArgs() Args {
	dsn := flag.String("dsn", "", "Database connection string")
	address := flag.String("address", "", "The server address to listen on")

	flag.Parse()

	if *dsn == "" {
		log.Fatalln("dsn flag is required")
	}

	if *address == "" {
		log.Fatalln("address flag is required")
	}

	return Args{
		Dsn:     *dsn,
		Address: *address,
	}
}
