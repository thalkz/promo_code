package endpoints

import (
	"fmt"
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
	var code promocode.Promocode
	err := c.BindJSON(&code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"reason": fmt.Sprintf("failed to parse request body: %v", err),
		})
		return
	}

	database.Instance[code.Name] = &code

	c.JSON(http.StatusOK, SuccessAddResponse{
		PromocodeName: code.Name,
		Status:        "added",
		Avantage:      code.Avantage,
	})
}
