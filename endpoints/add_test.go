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

type addTestCase struct {
	Request            string
	ExpectedHttpStatus int
	ExpectedResponse   endpoints.AddResponse
}

var addTestCases = []addTestCase{
	{
		Request: `{
			"_id": "WEATHER_CODE_ID",
			"name": "WeatherCode",
			"avantage": { "percent": 20 },
			"restrictions": [
			  {
				"@date": {
				  "after": "2019-01-01",
				  "before": "2020-06-30"
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
		  }`,
		ExpectedHttpStatus: http.StatusOK,
		ExpectedResponse: endpoints.AddResponse{
			PromocodeName: "WeatherCode",
			Status:        "accepted",
			Advantage: promocode.Avantage{
				Percent: 20,
			},
		},
	},
}

func TestHandleAdd(t *testing.T) {
	database.Reset()

	router := router.SetupRouter()

	for _, tc := range addTestCases {
		w := httptest.NewRecorder()

		bodyReader := bytes.NewReader([]byte(tc.Request))
		req, _ := http.NewRequest("PUT", "/add", bodyReader)

		router.ServeHTTP(w, req)
		assert.Equal(t, tc.ExpectedHttpStatus, w.Code)

		var response endpoints.AddResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "failed to unmarshall body")

		assert.Equal(t, response, tc.ExpectedResponse)

		assert.Equal(t, len(database.Instance), 1)
		promocode := database.Instance[tc.ExpectedResponse.PromocodeName]

		assert.Equal(t, promocode.Name, tc.ExpectedResponse.PromocodeName)
		// TODO test if the promocode is correct (json)
	}
}
