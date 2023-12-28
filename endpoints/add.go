package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/promocode"
)

func HandleAdd(c *gin.Context) {
	var request promocode.Promocode
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"promocode_name": request.Name,
			"status":         "bad request",
			"reason":         err.Error(),
		})
		return
	}

	database.Instance[request.Name] = &request
}
