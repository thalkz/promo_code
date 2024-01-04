package weather

type WeatherGetter interface {
	GetWeather(cityName string) (string, int, error)
}
