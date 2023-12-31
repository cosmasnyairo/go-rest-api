package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrorConnectingToDB       = "failed to connect to database"
	EmptyDatabase             = Database{}
	SuccessfullyConnectedToDB = "successfully connected to database"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"), os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE"),
	)
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &EmptyDatabase, fmt.Errorf("%v : %w", ErrorConnectingToDB, err)
	}
	fmt.Println(SuccessfullyConnectedToDB)

	return &Database{
		Client: dbConn,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	err := d.Client.DB.PingContext(ctx)
	return err
}
