package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adrianoff/aws/server/structures"
)

func ReadTemplate() []byte {
	template, err := os.ReadFile("../templates/index.html")
	if err != nil {
		fmt.Println(err)
	}

	return template
}

func PrepareHtml(template string, forecast *structures.OpenMeteoResponse) string {

	html := strings.Replace(template, "%current_weather_code%", strconv.Itoa(forecast.Current.WeatherCode), 1)
	html = strings.Replace(html, "%current_temperature%", fmt.Sprintf("%.1f", forecast.Current.Temperature), 1)
	html = strings.Replace(html, "%current_apparent_temperature%", fmt.Sprintf("%.1f", forecast.Current.ApparentTemperature), 1)
	html = strings.Replace(html, "%current_wind_direction_10m%", strconv.Itoa(forecast.Current.WindDirection), 1)
	html = strings.Replace(html, "%current_wind_speed_10m%", fmt.Sprintf("%.1f", forecast.Current.WindSpeed), 1)
	html = strings.Replace(html, "%daily_temperature_2m_max%", fmt.Sprintf("%.1f", forecast.Daily.TemperatureMax[0]), 1)
	html = strings.Replace(html, "%daily_temperature_2m_min%", fmt.Sprintf("%.1f", forecast.Daily.TemperatureMin[0]), 1)
	html = strings.Replace(html, "%daily_precipitation_sum%", fmt.Sprintf("%.1f", forecast.Daily.PrecipitationSum[0]), 1)
	html = strings.Replace(html, "%sunset%", forecast.Daily.Sunset[0], 1)
	html = strings.Replace(html, "%sunrise%", forecast.Daily.Sunrise[0], 1)

	html = strings.Replace(html, "\n", "", -1)

	return html
}