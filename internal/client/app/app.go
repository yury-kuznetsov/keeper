package app

import (
	"flag"
	"fmt"
	dataHandler "gophkeeper/internal/client/handler/data"
	userHandler "gophkeeper/internal/client/handler/user"
	"os"
)

// App структура приложения
type App struct {
	di *Container
}

// New creates a new instance of the App structure using the provided configuration. It initializes the dependency
// injection container and returns a pointer to the created App instance.
//
// Example:
//
//	config := map[string]string{
//	    "SERVER_ADDRESS": "https://example.com",
//	    "DB_FILE":        "/path/to/dbfile.db",
//	}
//	app := New(config)
//	app.Start()
func New(config map[string]string) *App {
	return &App{
		di: NewContainer(config),
	}
}

// Start executes the main logic of the application based on the provided command-line arguments.
// It checks if a subcommand is provided, and based on the subcommand, it calls the corresponding handler function.
// If an error occurs during the execution of the handler function, it prints the error message and exits the application.
func (app *App) Start() {
	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	var err error

	switch os.Args[1] {

	case "register":
		err = userHandler.RegisterHandler(app.di.UserSvc(), os.Args[2:])

	case "login":
		err = userHandler.LoginHandler(app.di.UserSvc(), os.Args[2:])

	case "create":
		err = dataHandler.CreateHandler(app.di.DataSvc(), os.Args[2:])

	case "update":
		err = dataHandler.UpdateHandler(app.di.DataSvc(), os.Args[2:])

	case "remove":
		err = dataHandler.RemoveHandler(app.di.DataSvc(), os.Args[2:])

	case "pull":
		err = dataHandler.PullHandler(app.di.DataSvc(), os.Args[2:])

	case "push":
		err = dataHandler.PushHandler(app.di.DataSvc(), os.Args[2:])

	case "status":
		err = dataHandler.StatusHandler(app.di.DataSvc())

	case "list":
		err = dataHandler.ListHandler(app.di.DataSvc())

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(1)
	}
}
