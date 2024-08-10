package structures

type Forecast struct {
	Current struct {
		WeatherCode string

		Temperature struct {
			Real     int
			Apparent int
			High     int
			Low      int
		}
		IsDay         bool
		Precipitation float32
		Wind          struct {
			Speed     int
			Direction int
		}
		Sunrise string
		Sunset  string
	}

	Hours [3]struct {
		WeatherCode string
		Visibility  int
		Humidity    int

		Temperature struct {
			Real     int
			Apparent int
		}
		Wind struct {
			Speed     int
			Direction int
		}
		Precipitation struct {
			Probability int
			Quantity    float32
		}
	}
}
