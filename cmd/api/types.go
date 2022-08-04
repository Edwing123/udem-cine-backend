package main

import (
	"flag"
	"log"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	SessionKey    = "session"
	UserIdKey     = "userId"
	IsLoggedInKey = "isLoggedIn"
)

type Api struct {
	Models      models.Models
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Store       *session.Store
}

func (api *Api) NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiber",
		AppName:      "UdeMCine API",
	})

	// Define global middlewares.
	app.Use(
		api.SetSessionToContext,
		recover.New(),
		logger.New(),
	)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:4173, http://localhost:5173",
	}))

	// No protected routes.
	app.Post("/auth/login", api.AuthLogin)
	app.Get("/is-logged-in", api.IsLoggedIn)

	// Protected routes.
	app.Use(api.AuthenticateRequest)
	app.Get("/user/:id", api.UserGet)
	app.Post("/auth/logout", api.AuthLogout)

	users := app.Group("/users")
	users.Get("/list", api.UsersList)
	users.Post("/create", api.UsersCreate)
	users.Patch("/edit", api.UsersEdit)
	users.Delete("/delete", api.UsersDelete)

	movies := app.Group("/movies")
	movies.Get("/get", api.MoviesGet)
	movies.Get("/list", api.MoviesList)
	movies.Post("/create", api.MoviesCreate)
	movies.Patch("/edit", api.MoviesEdit)
	movies.Delete("/delete", api.MoviesDelete)

	return app
}

// Response structs.
type ErrorMessage[T any] struct {
	Ok     bool `json:"ok"`
	Reason T    `json:"reason"`
}

type SuccessMessage[T any] struct {
	Ok   bool `json:"ok"`
	Data T    `json:"data"`
}

// Request bodies.
type ModelWithId[T any] struct {
	Id   int `json:"id"`
	Data T   `json:"data"`
}

type BodyWithId struct {
	Id int `json:"id"`
}

// Command line arguments.
type Args struct {
	Dsn           string
	Address       string
	StoreUserName string
	StorePassword string
	StoreDatabase string
}

func GetArgs() Args {
	dsn := flag.String("dsn", "", "Database connection string")

	address := flag.String("address", "", "The server address to listen on")

	storeUserName := flag.String(
		"storeUserName",
		"",
		"Username for the PostgreSQL sessions store",
	)

	storePassword := flag.String("storePassword", "", "Password for storeUserName")

	storeDatabase := flag.String("storeDatabase", "", "Name of the database for the sessions store")

	flag.Parse()

	if *dsn == "" {
		log.Fatalln("dsn flag is required")
	}

	if *address == "" {
		log.Fatalln("address flag is required")
	}

	if *storeUserName == "" {
		log.Fatalln("storeUserName flag is required")
	}

	if *storePassword == "" {
		log.Fatalln("storePassword flag is required")
	}

	if *storeDatabase == "" {
		log.Fatalln("storeDatabase flag is required")
	}

	return Args{
		Dsn:           *dsn,
		Address:       *address,
		StoreUserName: *storeUserName,
		StorePassword: *storePassword,
		StoreDatabase: *storeDatabase,
	}
}
