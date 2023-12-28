package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/promocode"
)

type verifyRequest struct {
	PromocodeName string `json:"promocode_name"`
	Arguments     promocode.Arguments
}

func HandleVerify(c *gin.Context) {
	var request verifyRequest
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"promocode_name": request.PromocodeName,
			"status":         "bad request",
			"reason":         err.Error(),
		})
		return
	}

	promocode, ok := database.Instance[request.PromocodeName]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"promocode_name": request.PromocodeName,
			"status":         "denied",
			"reason":         "promocode does not exist in database",
		})
		return
	}

	valid, err := promocode.Validate(request.Arguments)

	if valid {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"promocode_name": request.PromocodeName,
			"status":         "denied",
			"reason":         errors.Wrap(err, "validation failed"),
		})
	}
}
