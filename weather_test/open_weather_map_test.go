package weather_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/thalkz/promo_code/weather"
)

func TestHandleVerify(t *testing.T) {
	godotenv.Load("../.env")

	owm := weather.OpenWeatherMap{
		ApiKey: os.Getenv("OPEN_WEATHER_MAP_TEST_API_KEY"),
	}

	status, _, err := owm.GetWeather("Lyon")
	if err != nil {
		t.Fatalf("failed to get weather: %v", err)
	}
	if status == "" {
		t.Fatalf("invalid response status %v", status)
	}

	// TODO: Test if temperature is valid
}
