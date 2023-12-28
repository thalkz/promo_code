package endpoints_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/endpoints"
	"github.com/thalkz/promo_code/promocode"
	"github.com/thalkz/promo_code/router"
)

type verifyTestCase struct {
	Request            string
	ExpectedHttpStatus int
	ExpectedResponse   endpoints.VerifyResponse
}

var verifyTestCases = []verifyTestCase{
	{
		Request: `{
			"promocode_name": "WeatherCode",
			"arguments": {
			  "age": 40,
			  "meteo": { "town": "Lyon" }
			}
		  }`,
		ExpectedHttpStatus: http.StatusOK,
		ExpectedResponse: endpoints.VerifyResponse{
			PromocodeName: "WeatherCode",
			Status:        "accepted",
			Avantage: promocode.Avantage{
				Percent: 20,
			},
		},
	},
}

func generateTestPromocode() promocode.Promocode {
	var str = `{
		"_id": "WEATHER_CODE_ID",
		"name": "WeatherCode",
		"avantage": { "percent": 20 },
		"restrictions": [
		  {
			"@date": {
			  "after": "2019-01-01",
			  "before": "2024-06-30"
			}
		  },
		  {
			"@or": [
			  {
				"@age": {
				  "eq": 40
				}
			  },
			  {
				"@and": [
				  {
					"@age": {
					  "lt": 30,
					  "gt": 15
					}
				  },
				  {
					"@meteo": {
					  "is": "clear",
					  "temp": {
						"gt": "15"
					  }
					}
				  }
				]
			  }
			]
		  }
		]
	  }`

	var result promocode.Promocode
	json.Unmarshal([]byte(str), &result)
	return result
}

func TestHandleVerify(t *testing.T) {
	database.Reset()
	testPromocode := generateTestPromocode()
	database.Instance["WeatherCode"] = &testPromocode

	router := router.SetupRouter()

	for _, tc := range verifyTestCases {
		w := httptest.NewRecorder()

		bodyReader := bytes.NewReader([]byte(tc.Request))

		req, err := http.NewRequest("GET", "/verify", bodyReader)
		assert.NoError(t, err, "failed to created new request")

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var response endpoints.VerifyResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "failed to unmarshall body")

		assert.Equal(t, tc.ExpectedResponse, response)
	}
}
