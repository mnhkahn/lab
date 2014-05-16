package models

type Weather struct {
	Date    string          `json:"date"`
	Results []WeatherResult `json:"results"`
}

type WeatherResult struct {
	CurrentCity string        `json:"currentCity"`
	WeatherDate []WeatherDate `json:"weather_data"`
}

type WeatherDate struct {
	Date            string `json:"date"`
	PicUrl          string `json:"pic"`
	DayPictureUrl   string `json:"dayPictureUrl"`
	NightPictureUrl string `json:"nightPictureUrl"`
	Weather         string `json:"weather"`
	Wind            string `json:"wind"`
	Temperature     string `json:"temperature"`
}

type PM struct {
	PM2_5            string `json:"pm2_5"`
	Area             string `json:"area"`
	Pm2_5            string `json:"pm2_5"`
	Pm2_5_24h        string `json:"pm2_5_24h"`
	PositionName     string `json:"position_name"`
	PrimaryPollutant string `json:"primary_pollutant"`
	Quality          string `json:"quality"`
	StationCode      string `json:"station_code"`
	TimePoint        string `json:"time_point"`
}
