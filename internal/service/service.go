package service

import (
	"kwaka_test/internal/entity"
	"kwaka_test/internal/repository"
)

type Weather interface {
	CreateWeather(location string) (entity.Weather, error)
	GetWeather(location string) (entity.Weather, error)
	UpdateWeather(location string, input entity.UpdateWeather) error
}

type Service struct {
	Weather
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Weather: NewWeatherService(repo.Weather)}
}
