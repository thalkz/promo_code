package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OpenWeatherMap struct {
	ApiKey string
}

func (g OpenWeatherMap) GetWeather(cityName string) (string, int, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", cityName, g.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("query failed: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("failed to read response: %v", err)
	}

	var response openWeatherMapResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse response: %v", err)
	}

	if response.Code != http.StatusOK {
		return "", 0, fmt.Errorf("invalid response code, got %v (message: %v)", response.Code, response.Message)
	}

	if len(response.Weather) == 0 {
		return "", 0, fmt.Errorf("response.weather is empty in %v", string(body))
	}

	is := strings.ToLower(response.Weather[0].Main)
	temp := int(response.Main.Temp) // Floor is performed to get an int

	return is, temp, nil
}

type openWeatherMapResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Weather []openWeatherMapWeather `json:"weather"`
	Main    openWeatherMapMain      `json:"main"`
}

type openWeatherMapWeather struct {
	Main string `json:"main"`
}

type openWeatherMapMain struct {
	Temp float64 `json:"temp"`
}
