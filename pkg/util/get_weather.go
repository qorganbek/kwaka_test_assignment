package util

import (
	"encoding/json"
	"fmt"
	"kwaka_test/internal/entity"
	"net/http"
	"os"
)

func GetWeather(city string) (entity.Weather, error) {
	apiKey := os.Getenv("OPEN_API_WEATHER")
	apiUrl := fmt.Sprintf(os.Getenv("API_URL"), city, apiKey)

	response, err := http.Get(apiUrl)
	if err != nil {
		return entity.Weather{}, err
	}
	defer response.Body.Close()

	// Decode JSON response
	var weatherData struct {
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return entity.Weather{}, err
	}
	fmt.Println(weatherData)

	weather := entity.Weather{
		Location:    weatherData.Name,
		Temperature: weatherData.Main.Temp,
		Description: weatherData.Weather[0].Description,
		FeelsLike:   weatherData.Main.FeelsLike,
	}

	return weather, nil
}
