package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thalkz/promo_code/database"
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

func HandleVerify(c *gin.Context) {
	var request verifyRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "bad request",
			Reason:        err.Error(),
		})
		return
	}

	code, ok := database.Instance[request.PromocodeName]
	fmt.Println("Promocode found", code)

	if !ok {
		c.JSON(http.StatusOK, VerifyResponse{
			PromocodeName: request.PromocodeName,
			Status:        "denied",
			Reason:        "promocode does not exist in database",
		})
		return
	}

	var meteoStatus, meteoTemp = weather.Get(request.Arguments.Meteo.Town)

	argument := promocode.Arguments{
		Age:         request.Arguments.Age,
		Date:        time.Now(),
		MeteoStatus: meteoStatus,
		MeteoTemp:   meteoTemp,
	}

	valid, err := code.Validate(argument)

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
			Reason:        errors.Wrap(err, "validation failed").Error(),
		})
	}
}
