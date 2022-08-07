package main

import (
	"flag"
	"log"

	"github.com/Edwing123/udem-cine/pkg/codes"
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
		recover.New(),
		logger.New(),
		api.SetSessionToContext,
	)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	// No protected routes.
	app.Post("/auth/login", api.AuthLogin)
	app.Get("/is-logged-in", api.IsLoggedIn)

	// Protected routes.
	app.Use(api.OnlyAuthenticated)
	app.Post("/auth/logout", api.AuthLogout)

	users := app.Group("/users")
	users.Get("/get/:id", api.UsersGet)
	users.Use(api.OnlyAdmin)
	users.Get("/list", api.UsersList)
	users.Post("/create", api.UsersCreate)
	users.Patch("/edit/:id", api.UsersEdit)
	users.Delete("/delete/:id", api.UsersDelete)

	movies := app.Group("/movies", api.OnlyAdmin)
	movies.Get("/get/:id", api.MoviesGet)
	movies.Get("/list", api.MoviesList)
	movies.Post("/create", api.MoviesCreate)
	movies.Patch("/edit/:id", api.MoviesEdit)
	movies.Delete("/delete/:id", api.MoviesDelete)

	rooms := app.Group("/rooms", api.OnlyAdmin)
	rooms.Get("/get/:number", api.RoomsGet)
	rooms.Get("/list", api.RoomsList)
	rooms.Post("/create", api.RoomsCreate)
	rooms.Patch("/edit/:number", api.RoomsEdit)
	rooms.Delete("/delete/:number", api.RoomsDelete)

	schedules := app.Group("/schedules", api.OnlyAdmin)
	schedules.Get("/get/:id", api.SchedulesGet)
	schedules.Get("/list", api.SchedulesList)
	schedules.Post("/create", api.SchedulesCreate)
	schedules.Patch("/edit/:id", api.SchedulesEdit)
	schedules.Delete("/delete/:id", api.SchedulesDelete)

	functions := app.Group("/functions")
	functions.Get("/list", api.FunctionsList)
	functions.Use(api.OnlyAdmin)
	functions.Get("/get/:id", api.FunctionsGet)
	functions.Post("/create", api.FunctionsCreate)
	functions.Patch("/edit/:id", api.FunctionsEdit)
	functions.Delete("/delete/:id", api.FunctionsDelete)

	return app
}

// Response structs.
type ErrorMessage[T any] struct {
	Ok      bool       `json:"ok"`
	Code    codes.Code `json:"code"`
	Details T          `json:"details,omitempty"`
}

type SuccessMessage[T any] struct {
	Ok   bool `json:"ok"`
	Data T    `json:"data"`
}

// Command line arguments.
type Args struct {
	Dsn           string
	Address       string
	StoreUserName string
	StorePassword string
	StoreDatabase string
	CertKeyPath   string
	CertPath      string
}

func GetArgs() Args {
	dsn := flag.String("dsn", "", "Database connection string")

	address := flag.String("address", "", "The server address to listen on")

	certPath := flag.String("certPath", "", "Path of the TLS cert")

	certKeyPath := flag.String("certKeyPath", "", "Path of the TLS cert key")

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

	if *certPath == "" {
		log.Fatalln("certPath flag is required")
	}

	if *certKeyPath == "" {
		log.Fatalln("certKeyPath flag is required")
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
		CertPath:      *certPath,
		CertKeyPath:   *certKeyPath,
	}
}
