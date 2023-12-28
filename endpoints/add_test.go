package endpoints_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/router"
)

type addTestCase struct {
	Request        string
	ExpectedStatus string
	ExpectedName   string
}

var addTestCases = []addTestCase{}

func TestHandleAdd(t *testing.T) {
	database.Reset()

	router := router.SetupRouter()

	for _, tc := range addTestCases {
		w := httptest.NewRecorder()

		jsonBody := []byte(tc.Request)
		bodyReader := bytes.NewReader(jsonBody)

		req, err := http.NewRequest("GET", "/add", bodyReader)
		assert.NoError(t, err, "failed to created new request")

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var response map[string]string
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "failed to unmarshall body")

		assert.Equal(t, response["promocode_name"], tc.ExpectedName)
		assert.Equal(t, response["status"], tc.ExpectedStatus)

		assert.Equal(t, len(database.Instance), 1)
		promocode := database.Instance[tc.ExpectedName]

		assert.Equal(t, promocode.Name, tc.ExpectedName)
		// TODO test if the promocode is correct (json)
	}
}
