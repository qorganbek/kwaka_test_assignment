package repository

import (
	"github.com/jmoiron/sqlx"
)

type WeatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgres(db *sqlx.DB) *WeatherPostgres {
	return &WeatherPostgres{db: db}
}
