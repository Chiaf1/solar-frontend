package repositories

type WeatherRepository interface {
	GetCurrentTemperature(lat, lon float64) (float64, error)
}
