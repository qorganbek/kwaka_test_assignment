package repository

import (
	"fmt"
	"kwaka_test/internal/entity"
	"kwaka_test/internal/repository/pgrepo"
	"log"
)

func (w WeatherPostgres) CreateWeather(weather entity.Weather) (int, error) {
	tableCreationQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
													id          SERIAL PRIMARY KEY,
													location    VARCHAR(255) unique,
													description VARCHAR(255),
													temp        float,
													feels_like float);`,
		pgrepo.WeatherTable)
	_, err := w.db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (location, description, temp, feels_like) VALUES ($1, $2, $3, $4) RETURNING id", pgrepo.WeatherTable)

	row := w.db.QueryRow(query, weather.Location, weather.Description, weather.Temperature, weather.FeelsLike)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (w WeatherPostgres) GetWeather(location string) (entity.Weather, error) {
	var weather entity.Weather
	query := fmt.Sprintf("SELECT * FROM %s WHERE location=$1", pgrepo.WeatherTable)
	err := w.db.Select(&weather, query, location)
	return weather, err
}

func (w WeatherPostgres) UpdateWeather(location string, input entity.UpdateWeather) error {
	query := fmt.Sprintf("UPDATE %s SET temp=$1, description=$2, feels_like=$3 WHERE location=$4", pgrepo.WeatherTable)
	_, err := w.db.Exec(query, input.Temperature, input.Description, input.FeelsLike, location)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
