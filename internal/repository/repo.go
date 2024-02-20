package repository

import (
	"github.com/jmoiron/sqlx"
	"kwaka_test/internal/entity"
)

type Weather interface {
	CreateWeather(location string) (entity.Weather, error)
	GetWeather(location string) (entity.Weather, error)
	UpdateWeather(location string, input entity.UpdateWeather) error
}

type Repository struct {
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Weather: NewWeatherPostgres(db),
	}
}
