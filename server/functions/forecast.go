package functions

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/adrianoff/aws/server/structures"
)

func GetOpenMeteoForecast() (*structures.OpenMeteoResponse, error) {
	var openMeteoResponse structures.OpenMeteoResponse
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=44.92938294579323&longitude=34.08182283665256&current=temperature_2m,apparent_temperature,is_day,precipitation,rain,showers,snowfall,weather_code,wind_speed_10m,wind_direction_10m&hourly=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,precipitation,rain,weather_code,visibility,wind_speed_10m,wind_direction_10m&daily=temperature_2m_max,temperature_2m_min,sunrise,sunset,precipitation_sum,precipitation_probability_max&timezone=Europe%2FMoscow&forecast_days=1")
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
