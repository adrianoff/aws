package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adrianoff/aws/server/structures"
)

var weatherCodeMap = map[int]string{
	0: "Clear sky",
	1: "Mainly clear",
	2: "Partly cloudy",
	3: "Overcast",

	45: "Fog",
	48: "Rime",

	51: "Drizzle",
	53: "Drizzle",
	55: "Drizzle",
	56: "Freezing Drizzle",
	57: "Freezing Drizzle",

	61: "Slight rain",
	63: "Rain",
	65: "Heavy rain",
	66: "Freezing rain",
	67: "Freezing rain",

	71: "Slight snow",
	73: "Snow",
	75: "Heavy snow",
	77: "Snow grains",

	80: "Slight shower",
	81: "Shower",
	82: "Heavy shower",
	85: "Snow shower",
	86: "Snow shower",

	95: "Thunderstorm",
	96: "Thunderstorm, hail",
	99: "Thunderstorm, heavy hail",
}

var weatherCodeIcons = map[string]string{
	"Clear sky day":       "fa-sun",
	"Mainly clear day":    "fa-sun",
	"Partly cloudy day":   "fa-cloud-sun",
	"Overcast day":        "fa-cloud",
	"Clear sky night":     "fa-moon",
	"Mainly clear night":  "fa-moon",
	"Partly cloudy night": "fa-cloud-moon",
	"Overcast night":      "fa-cloud",

	"Fog day":    "fa-smog",
	"Rime day":   "fa-smog",
	"Fog night":  "fa-smog",
	"Rime night": "fa-smog",

	"Drizzle day":            "fa-cloud-rain",
	"Freezing Drizzle day":   "fa-cloud-rain",
	"Drizzle night":          "fa-cloud-rain",
	"Freezing Drizzle night": "fa-cloud-rain",

	"Slight rain day":     "fa-cloud-sun-rain",
	"Rain day":            "fa-cloud-rain",
	"Heavy rain day":      "fa-cloud-showers-heavy",
	"Freezing rain day":   "fa-cloud-rain",
	"Slight rain night":   "fa-cloud-moon-rain",
	"Rain night":          "fa-cloud-rain",
	"Heavy rain night":    "fa-cloud-showers-heavy",
	"Freezing rain night": "fa-cloud-rain",

	"Slight snow day":   "fa-snowflake",
	"Snow day":          "fa-snowflake",
	"Heavy snow day":    "fa-snowflake",
	"Snow grains day":   "fa-snowflake",
	"Slight snow night": "fa-snowflake",
	"Snow night":        "fa-snowflake",
	"Heavy snow night":  "fa-snowflake",
	"Snow grains night": "fa-snowflake",

	"Slight shower day":   "fa-cloud-showers-heavy",
	"Shower day":          "fa-cloud-showers-heavy",
	"Heavy shower day":    "fa-cloud-showers-heavy",
	"Snow shower day":     "fa-cloud-showers-heavy",
	"Slight shower night": "fa-cloud-showers-heavy",
	"Shower night":        "fa-cloud-showers-heavy",
	"Heavy shower night":  "fa-cloud-showers-heavy",
	"Snow shower night":   "fa-cloud-showers-heavy",

	"Thunderstorm day":               "fa-cloud-bolt",
	"Thunderstorm, hail day":         "fa-cloud-bolt",
	"Thunderstorm, heavy hail day":   "fa-cloud-bolt",
	"Thunderstorm night":             "fa-cloud-bolt",
	"Thunderstorm, hail night":       "fa-cloud-bolt",
	"Thunderstorm, heavy hail night": "fa-cloud-bolt",
}

var isDayMap = map[int]string{
	0: "night",
	1: "day",
}

func ReadTemplate() []byte {
	template, err := os.ReadFile("../templates/index.html")
	if err != nil {
		fmt.Println(err)
	}

	return template
}

func PrepareHtml(template string, forecast *structures.OpenMeteoResponse) string {

	parseTime, err := time.Parse("2006-01-02T15:04", forecast.Current.Time)
	if err != nil {
		panic(err)
	}

	dateStr := parseTime.Format("2 Jan 2006")
	dayOfWeek := parseTime.Weekday().String()
	timeStr := parseTime.Format("15:04")

	weatherCode, ok := weatherCodeMap[forecast.Current.WeatherCode]
	if !ok {
		weatherCode = "N/A"
	}

	dayOrNightStr, ok := isDayMap[forecast.Current.IsDay]
	if !ok {
		dayOrNightStr = "day"
	}
	iconStr, ok := weatherCodeIcons[weatherCode+" "+dayOrNightStr]
	if !ok {
		iconStr = "fa-circle-question"
	}

	parseTimeSunset, err := time.Parse("2006-01-02T15:04", forecast.Daily.Sunset[0])
	if err != nil {
		panic(err)
	}
	sunsetStr := parseTimeSunset.Format("15:04")
	parseTimeSunrise, err := time.Parse("2006-01-02T15:04", forecast.Daily.Sunrise[0])
	if err != nil {
		panic(err)
	}
	sunriseStr := parseTimeSunrise.Format("15:04")

	html := strings.Replace(template, "%current_date%", dateStr, 1)
	html = strings.Replace(html, "%current_dayofweek%", dayOfWeek, 1)
	html = strings.Replace(html, "%current_time%", timeStr, 1)
	html = strings.Replace(html, "%current_weather_code%", weatherCode, 1)
	html = strings.Replace(html, "%current_weather_code_icon%", iconStr, 1)
	html = strings.Replace(html, "%current_temperature%", fmt.Sprintf("%.1f", forecast.Current.Temperature), 1)
	html = strings.Replace(html, "%current_apparent_temperature%", fmt.Sprintf("%.1f", forecast.Current.ApparentTemperature), 1)
	html = strings.Replace(html, "%current_humidity%", strconv.Itoa(forecast.Current.Humidity), 1)
	html = strings.Replace(html, "%current_pressure%", fmt.Sprintf("%.1f", forecast.Current.Pressure), 1)
	html = strings.Replace(html, "%current_wind_direction_10m%", strconv.Itoa(forecast.Current.WindDirection), 1)
	html = strings.Replace(html, "%current_wind_speed_10m%", fmt.Sprintf("%.1f", forecast.Current.WindSpeed), 1)
	html = strings.Replace(html, "%daily_temperature_2m_max%", fmt.Sprintf("%.1f", forecast.Daily.TemperatureMax[0]), 1)
	html = strings.Replace(html, "%daily_temperature_2m_min%", fmt.Sprintf("%.1f", forecast.Daily.TemperatureMin[0]), 1)
	html = strings.Replace(html, "%daily_precipitation_sum%", fmt.Sprintf("%.1f", forecast.Daily.PrecipitationSum[0]), 1)
	html = strings.Replace(html, "%sunset%", sunsetStr, 1)
	html = strings.Replace(html, "%sunrise%", sunriseStr, 1)


	currentHour := parseTime.Hour()
	for i := 1; i <= 3; i++ {
		hour := currentHour+i+1
		hourStr := "hour"+strconv.Itoa(i)

		hourParseTime, err := time.Parse("2006-01-02T15:04", forecast.Hourly.Time[hour])
		if err != nil {
			panic(err)
		}
		hourDayOrNightStr, ok := isDayMap[forecast.Hourly.IsDay[hour]]
		if !ok {
			hourDayOrNightStr = "day"
		}
		hourWeatherCode, ok := weatherCodeMap[forecast.Hourly.WeatherCode[hour]]
		if !ok {
			hourWeatherCode = "N/A"
		}
		hourWeatherIcon, ok := weatherCodeIcons[hourWeatherCode+" "+hourDayOrNightStr]
		if !ok {
			hourWeatherIcon = "fa-circle-question"
		}

		html = strings.Replace(html, "%" + hourStr + "_time%", hourParseTime.Format("15:04"), 1)
		html = strings.Replace(html, "%" + hourStr + "_weather_code%", hourWeatherCode, 1)
		html = strings.Replace(html, "%" + hourStr + "_weather_code_icon%", hourWeatherIcon, 1)
		html = strings.Replace(html, "%" + hourStr + "_temperature%", fmt.Sprintf("%.1f", forecast.Hourly.Temperature[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_apparent_temperature%", fmt.Sprintf("%.1f", forecast.Hourly.ApparentTemperature[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_wind_direction_10m%", strconv.Itoa(forecast.Hourly.WindDirection[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_wind_speed_10m%", fmt.Sprintf("%.1f", forecast.Hourly.WindSpeed[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_precipitation%", fmt.Sprintf("%.1f", forecast.Hourly.Precipitation[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_precipitation_probability%", strconv.Itoa(forecast.Hourly.PrecipitationProbability[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_visibility%", fmt.Sprintf("%.1f", forecast.Hourly.Visibility[hour]), 1)
		html = strings.Replace(html, "%" + hourStr + "_humidity%", strconv.Itoa(forecast.Hourly.Humidity[hour]), 1)
	}

	html = strings.Replace(html, "\n", "", -1)

	return html
}
