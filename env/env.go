package env

import "os"

func GetOpenWeatherApiKey() string {
	return os.Getenv("OPEN_WEATHER_API_KEY")
}
