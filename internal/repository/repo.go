package repository

import (
	"github.com/jmoiron/sqlx"
	"kwaka_test/internal/entity"
)

type WeatherRepository interface {
	CreateWeather(weather entity.Weather) (int, error)
	GetWeather(location string) (entity.Weather, error)
	UpdateWeather(location string, input entity.UpdateWeather) error
}

type Repository struct {
	WeatherRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		WeatherRepository: NewWeatherPostgres(db),
	}
}
