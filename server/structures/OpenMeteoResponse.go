package structures

type OpenMeteoResponse struct {
	Latitude         float32 `json:"latitude"`
	Longitude        float32 `json:"longitude"`
	UtcOffsetSeconds int     `json:"utc_offset_seconds"`
	Current          struct {
		Time                string  `json:"time"`
		Temperature         float32 `json:"temperature_2m"`
		ApparentTemperature float32 `json:"apparent_temperature"`
		IsDay               int     `json:"is_day"`
		Precipitation       float32 `json:"precipitation"`
		WindSpeed           float32 `json:"wind_speed_10m"`
		WindDirection       int     `json:"wind_direction_10m"`
		WeatherCode         int     `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		TemperatureMax              []float32 `json:"temperature_2m_max"`
		TemperatureMix              []float32 `json:"temperature_2m_min"`
		Sunrise                     []string  `json:"sunrise"`
		Sunset                      []string  `json:"sunset"`
		PrecipitationSum            []float32 `json:"precipitation_sum"`
		PrecipitationProbabilityMax []float32 `json:"precipitation_probability_max"`
	} `json:"daily"`
	Hourly struct {
		Time                     []string  `json:"time"`
		Temperature              []float32 `json:"temperature_2m"`
		ApparentTemperature      []float32 `json:"apparent_temperature"`
		PrecipitationProbability []int     `json:"precipitation_probability"`
		Precipitation            []float32 `json:"precipitation"`
		WeatherCode              []int     `json:"weather_code"`
		Visibility               []int     `json:"visibility"`
		WindSpeed                []float32 `json:"wind_speed_10m"`
		WindDirection            []int     `json:"wind_direction_10m"`
	} `json:"hourly"`
}
