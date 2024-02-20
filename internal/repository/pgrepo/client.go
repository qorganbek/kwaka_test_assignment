package pgrepo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Port     string
	Host     string
	DBName   string
	Username string
	Password string
	SSLMode  string
}

const (
	WeatherTable = "weather"
)

func NewPostgresDB(pg Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			pg.Host, pg.Port, pg.DBName, pg.Username, pg.Password, pg.SSLMode))

	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
