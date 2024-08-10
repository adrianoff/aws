package structures

type OpenMeteoResponse struct {
	Latitude float32 `json:"latitude"`
	Current  struct {
		Temperature float32 `json:"temperature_2m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
	Hourly struct {
		Time        []string  `json:"time"`
		Temperature []float32 `json:"temperature_2m"`
	} `json:"hourly"`
}
