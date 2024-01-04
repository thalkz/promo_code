package weather

type WeatherGetter interface {
	GetWeather(cityName string, apiKey string) (string, int, error)
}
