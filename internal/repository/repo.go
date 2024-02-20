package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/qorganbek/kwaka_test_assignment/internal/entity"
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
