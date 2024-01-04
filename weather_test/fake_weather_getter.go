package weather_test

import "fmt"

type SuccessFakeWeatherGetter struct{}

func (g SuccessFakeWeatherGetter) GetWeather(cityName string) (string, int, error) {
	return "clear", 15, nil
}

type FailFakeWeatherGetter struct{}

func (g FailFakeWeatherGetter) GetWeather(cityName string) (string, int, error) {
	return "", 0, fmt.Errorf("this fake weather getter always fails")
}
