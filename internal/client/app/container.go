package app

import (
	"database/sql"
	"gophkeeper/internal/client/api"
	dataRepo "gophkeeper/internal/client/repository/data"
	dataSvc "gophkeeper/internal/client/service/data"
	userSvc "gophkeeper/internal/client/service/user"
	"log"

	// sql driver
	_ "github.com/mattn/go-sqlite3"
)

// Container is a struct that represents a container for various components and services.
//
// It has the following fields:
// - config: a map[string]string for storing configuration values.
// - client: a *api.Client for making API requests.
// - db: a *sql.DB for database operations.
// - dataRepo: a *dataRepo.Repository for interacting with data repository.
// - dataSvc: a *dataSvc.Service for data-related operations.
// - userSvc: a *userSvc.Service for user-related operations.
type Container struct {
	config   map[string]string
	client   *api.Client
	db       *sql.DB
	dataRepo *dataRepo.Repository
	dataSvc  *dataSvc.Service
	userSvc  *userSvc.Service
}

// NewContainer returns a new instance of Container with the provided configuration.
// The Container struct holds various components and services.
// The config parameter is a map[string]string for storing configuration values.
// The returned value is a pointer to the created Container instance.
func NewContainer(config map[string]string) *Container {
	return &Container{config: config}
}

// Client returns the API client instance.
//
// It initializes the api.Client instance if it is not yet initialized.
// The api.Client instance is created with the server address from the configuration map.
// The returned value is a pointer to the api.Client instance.
func (c *Container) Client() *api.Client {
	if c.client == nil {
		c.client = api.NewClient(c.config["SERVER_ADDRESS"])
	}

	return c.client
}

// DB returns the *sql.DB instance for database operations.
//
// It initializes the *sql.DB instance if it is not yet initialized.
// The *sql.DB instance is created with the "sqlite3" driver and the database file path from the configuration map.
// If the connection to the database fails, it logs the error and exits the program.
// The returned value is a pointer to the *sql.DB instance.
func (c *Container) DB() *sql.DB {
	if c.db == nil {
		db, err := sql.Open("sqlite3", c.config["DB_FILE"])
		if err != nil {
			log.Fatal(err)
		}
		c.db = db
	}

	return c.db
}

// DataRepo returns the data repository instance.
//
// It initializes the dataRepo.Repository instance if it is not yet initialized.
// The dataRepo.Repository instance is created with the Database instance obtained
// through the DB() method.
// The returned value is a pointer to the dataRepo.Repository instance.
func (c *Container) DataRepo() *dataRepo.Repository {
	if c.dataRepo == nil {
		c.dataRepo = dataRepo.NewRepository(c.DB())
	}

	return c.dataRepo
}

// DataSvc returns the data service instance.
//
// It initializes the dataSvc.Service instance if it is not yet initialized.
// The dataSvc.Service instance is created with the dataRepo.Repository instance obtained through the DataRepo() method
// and the api.Client instance obtained through the Client() method.
// The returned value is a pointer to the dataSvc.Service instance.
func (c *Container) DataSvc() *dataSvc.Service {
	if c.dataSvc == nil {
		c.dataSvc = dataSvc.NewService(c.DataRepo(), c.Client())
	}

	return c.dataSvc
}

// UserSvc returns the user service instance.
//
// It initializes the userSvc.Service instance if it is not yet initialized.
// The userSvc.Service instance is created with the api.Client instance obtained through the Client() method.
// The returned value is a pointer to the userSvc.Service instance.
func (c *Container) UserSvc() *userSvc.Service {
	if c.userSvc == nil {
		c.userSvc = userSvc.NewService(c.Client())
	}

	return c.userSvc
}
