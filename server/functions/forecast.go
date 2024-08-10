package functions

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/adrianoff/aws/server/structures"
)

func GetOpenMeteoForecast() (*structures.OpenMeteoResponse, error) {
	var openMeteoResponse structures.OpenMeteoResponse
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation,weather_code,wind_speed_10m,wind_direction_10m&hourly=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,precipitation,weather_code,visibility,wind_speed_10m,wind_direction_10m&daily=weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset,precipitation_sum,precipitation_probability_max,wind_speed_10m_max&timezone=Europe%2FMoscow")
	if err != nil {
		return &openMeteoResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &openMeteoResponse); err != nil {
		return &openMeteoResponse, err
	}

	return &openMeteoResponse, nil
}
