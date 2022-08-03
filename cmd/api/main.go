package main

import "github.com/gofiber/storage/postgres"

func main() {
	args := GetArgs()

	models := NewDatabase(args.Dsn)

	store := NewStore(postgres.Config{
		Username: args.StoreUserName,
		Password: args.StorePassword,
		Database: args.StoreDatabase,
		Table:    "gofiber_store",
	})

	api := Api{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		Models:      models,
		Store:       store,
	}

	app := api.NewApp()

	errorLogger.Fatalln(app.Listen(args.Address))
}
