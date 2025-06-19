type WeatherClient interface {
	GetDailyForecast(location metoc.Location, date metoc.Dtg) metoc.DailyForecast
}
