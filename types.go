package weather

// Response
type (
	ForecastResponse struct {
		Location Location `json:"location"`
		Forecast Forecast `json:"forecast"`
	}

	Location struct {
		Name string `json:"name"`
	}

	Forecast struct {
		Forecastday []Day `json:"forecastday"`
	}

	Day struct {
		Date  string `json:"date"`
		Hours []Hour `json:"hour"`
	}

	Hour struct {
		Time          string    `json:"time"`
		Temp          float64   `json:"temp_c"`
		TempFeelsLike float64   `json:"feelslike_c"`
		Condition     Condition `json:"condition"`
		WindSpeed     float64   `json:"wind_kph"`
		WindDir       string    `json:"wind_dir"`
		ChanceOfRain  float64   `json:"chance_of_rain"`
	}

	Condition struct {
		Text string `json:"text"`
	}
)
