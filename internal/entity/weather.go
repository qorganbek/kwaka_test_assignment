package entity

type Weather struct {
	ID          int     `json:"id"`
	Location    string  `json:"location" db:"location"`
	Temperature float64 `json:"temp" db:"temp"`
	Description string  `json:"description" db:"description"`
	FeelsLike   float64 `json:"feels_like" db:"feels_like"`
}

type UpdateWeather struct {
	Temperature float64 `json:"temp"`
	Description string  `json:"description"`
	FeelsLike   float64 `json:"feels_like"`
}
