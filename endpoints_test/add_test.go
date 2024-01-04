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
	"github.com/thalkz/promo_code/test_helper"
)

type addTestCase struct {
	Request            string
	ExpectedHttpStatus int
	ExpectedResponse   endpoints.SuccessAddResponse
}

var addTestCases = []addTestCase{
	{
		Request: `{
			"_id": "TEST_ID",
			"name": "TestCode",
			"avantage": { "percent": 10 },
			"restrictions": [
			  {
				"@meteo": {
				  "is": "foggy",
				  "temp": {
					"eq": "30"
				  }
				}
			  }
			]
		  }`,
		ExpectedHttpStatus: http.StatusOK,
		ExpectedResponse: endpoints.SuccessAddResponse{
			PromocodeName: "TestCode",
			Status:        "added",
			Avantage: promocode.Avantage{
				Percent: 10,
			},
		},
	},
	{
		Request: `{
			"_id": "TEST_ID",
			"name": "TestCode",
			"avantage": { "percent": 10 },
			"restrictions": [
			  {
				"@meteo": {
				  "is": "fo`,
		ExpectedHttpStatus: http.StatusBadRequest,
		ExpectedResponse: endpoints.SuccessAddResponse{
			Status: "bad request",
		},
	},
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
		ExpectedResponse: endpoints.SuccessAddResponse{
			PromocodeName: "WeatherCode",
			Status:        "added",
			Avantage: promocode.Avantage{
				Percent: 20,
			},
		},
	},
}

func TestHandleAdd(t *testing.T) {
	router := router.SetupRouter()

	for caseId, tc := range addTestCases {
		// Reset the database (since it's a global variable)
		database.Reset()

		// Make the HTTP request
		w := httptest.NewRecorder()
		bodyReader := bytes.NewReader([]byte(tc.Request))
		req, _ := http.NewRequest("PUT", "/add", bodyReader)
		router.ServeHTTP(w, req)

		// Verify returned status code
		assert.Equal(t, tc.ExpectedHttpStatus, w.Code)

		// Verify returned body
		var response endpoints.SuccessAddResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "failed to unmarshall body")
		assert.Equal(t, tc.ExpectedResponse, response)

		if tc.ExpectedResponse.Status == "added" {
			// Verify database has been updated with correct value
			assert.Equal(t, 1, len(database.Instance))
			var expected promocode.Promocode
			err = json.Unmarshal([]byte(tc.Request), &expected)
			if err == nil {
				code := database.Instance[tc.ExpectedResponse.PromocodeName]
				test_helper.AssertSameJson(t, caseId, expected, code)
			}
		} else {
			// Expected response is an error; Database should bne empty
			assert.Equal(t, 0, len(database.Instance))
		}
	}
}
