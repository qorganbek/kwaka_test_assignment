package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Weather represents the weather data structure
type Weather struct {
	Location    string  `json:"name"`
	Temperature float64 `json:"temp"`
	Description string  `json:"description"`
	FeelsLike   float64 `json:"feels_like"`
}

// GetWeather retrieves weather data from an API
func GetWeather(city string) (Weather, error) {
	// Replace "YOUR_API_KEY" with your actual API key
	apiKey := "b7ffe634c060fb5ad1da59af78646989"
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	// Send GET request to the API
	response, err := http.Get(apiUrl)
	if err != nil {
		return Weather{}, err
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
		return Weather{}, err
	}
	fmt.Println(weatherData)
	// Populate Weather struct
	weather := Weather{
		Location:    weatherData.Name,
		Temperature: weatherData.Main.Temp,
		Description: weatherData.Weather[0].Description,
		FeelsLike:   weatherData.Main.FeelsLike,
	}

	return weather, nil
}

func main() {
	// Example usage
	city := "Almaty"
	weather, err := GetWeather(city)
	if err != nil {
		fmt.Println("Error getting weather:", err)
		return
	}

	fmt.Printf("Weather in %s:\n", weather.Location)
	fmt.Printf("Temperature: %.2f°C\n", weather.Temperature)
	fmt.Printf("Description: %s\n", weather.Description)
	fmt.Printf("Temperature feels like: %.2f°C\n", weather.FeelsLike)
}
