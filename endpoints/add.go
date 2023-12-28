package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/promocode"
)

type SuccessAddResponse struct {
	PromocodeName string             `json:"promocode_name"`
	Status        string             `json:"status"`
	Avantage      promocode.Avantage `json:"avantage"`
}

func HandleAdd(c *gin.Context) {
	var request promocode.Promocode
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"promocode_name": request.Name,
			"status":         "bad request",
			"reason":         err.Error(),
		})
		return
	}

	database.Instance[request.Name] = &request

	c.JSON(http.StatusOK, SuccessAddResponse{
		PromocodeName: request.Name,
		Status:        "added",
		Avantage:      request.Avantage,
	})
}
