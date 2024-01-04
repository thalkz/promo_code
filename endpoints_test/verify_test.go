package endpoints_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func TestHandleVerify(t *testing.T) {
	setupTestDatabase()    // Initialize database with test values
	setupNow("2023-12-28") // Stub time.Now
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

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
