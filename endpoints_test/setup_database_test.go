package endpoints_test

import (
	"encoding/json"

	"github.com/thalkz/promo_code/database"
	"github.com/thalkz/promo_code/promocode"
)

var weatherCodeStr = `{
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

var testCodeStr = `{
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
  }`

var invalidCodeStr = `{
	"_id": "TEST_ID",
	"name": "TestCode",
	"avantage": { "percent": 10 },
	"restrictions": [
	  {
		"@meteo": {
		  "is": "fog`

func setupTestDatabase() {
	var weatherCode promocode.Promocode
	json.Unmarshal([]byte(weatherCodeStr), &weatherCode)

	var testCode promocode.Promocode
	json.Unmarshal([]byte(testCodeStr), &testCode)

	database.Reset()
	database.Instance["WeatherCode"] = &weatherCode
	database.Instance["TestCode"] = &testCode
}
