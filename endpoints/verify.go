package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/env"
	"github.com/thalkz/promo_code/promocode"
	"github.com/thalkz/promo_code/weather"
)

type verifyRequest struct {
	PromocodeName string                 `json:"promocode_name"`
	Arguments     verifyRequestArguments `json:"arguments"`
}

type verifyRequestArguments struct {
	Age   int               `json:"age"`
	Meteo meteoTownArgument `json:"meteo"`
}

type meteoTownArgument struct {
	Town string `json:"town"`
}

type VerifyResponse struct {
	PromocodeName string             `json:"promocode_name"`
	Status        string             `json:"status"`
	Avantage      promocode.Avantage `json:"avantage,omitempty"`
	Reason        string             `json:"reason,omitempty"`
}

// Thises functions are stubbed in tests to use mock data
var Now = time.Now
var GetWeatherApi = func() weather.WeatherGetter {
	apiKey := env.GetOpenWeatherApiKey()
	return weather.OpenWeatherMap{
		ApiKey: apiKey,
	}
}

func HandleVerify(c *gin.Context) {
	var request verifyRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "error",
			Reason:        fmt.Sprintf("failed to parse request: %v", err),
		})
		return
	}

	code, ok := database.Instance[request.PromocodeName]

	if !ok {
		c.JSON(http.StatusOK, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "denied",
			Reason:        "promocode does not exist in database",
		})
		return
	}

	weatherApi := GetWeatherApi()

	meteoStatus, meteoTemp, err := weatherApi.GetWeather(request.Arguments.Meteo.Town)
	if err != nil {
		c.JSON(http.StatusOK, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "error",
			Reason:        fmt.Sprintf("failed to get weather data: %v", err),
		})
		return
	}

	args := promocode.Arguments{
		Age:         request.Arguments.Age,
		Date:        Now(),
		MeteoStatus: meteoStatus,
		MeteoTemp:   meteoTemp,
	}

	valid, err := code.Validate(args)

	if valid {
		c.JSON(http.StatusOK, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "accepted",
			Avantage:      code.Avantage,
		})
	} else {
		c.JSON(http.StatusOK, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "denied",
			Reason:        fmt.Sprintf("promocode validation failed: %v", err),
		})
	}
}
