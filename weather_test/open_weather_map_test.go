package weather_test

import (
	"os"
	"testing"

	"github.com/thalkz/promo_code/weather"
)

func TestHandleVerify(t *testing.T) {
	owm := weather.OpenWeatherMap{
		ApiKey: os.Getenv("OPEN_WEATHER_MAP_TEST_API_KEY"),
	}

	status, temp, err := owm.GetWeather("Lyon")
	if err != nil {
		t.Fatalf("failed to get weather: %v", err)
	}
	if status == "" || temp == 0 {
		t.Fatalf("invalid response status %v or temp %v", status, temp)
	}
}
