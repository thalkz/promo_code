package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/promocode"
)

type AddResponse struct {
	PromocodeName string             `json:"promocode_name"`
	Status        string             `json:"status"`
	Advantage     promocode.Avantage `json:"advantage"`
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

	c.JSON(http.StatusOK, AddResponse{
		PromocodeName: request.Name,
		Status:        "accepted",
		Advantage:     request.Avantage,
	})
}
