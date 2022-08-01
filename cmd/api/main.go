package main

func main() {
	args := GetArgs()

	models := NewDatabase(args.Dsn)

	api := Api{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		Models:      models,
	}

	app := api.NewApp()

	errorLogger.Fatalln(app.Listen(args.Address))
}
