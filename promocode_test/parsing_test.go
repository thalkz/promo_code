package promocode_test

import (
	"encoding/json"
	"testing"

	"github.com/thalkz/promo_code/promocode"
)

type parsingTestCase struct {
	Json     string
	Expected promocode.Validator
}

func testParsing(t *testing.T, caseId int, tc parsingTestCase, actual any, err error) {
	if tc.Expected == nil {
		// When Expected is nil, we want the parsing to fail and return an error
		if err == nil {
			t.Errorf("TestCase #%v: expected parsing to fail, but no error was returned", caseId)
		}
	} else {
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", caseId, err)
		}
		assertSameJson(t, caseId, tc.Expected, actual)
	}
}

func TestAgeParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"eq": 30
			}`,
			Expected: promocode.NewAgeRestriction(ptr(30), nil, nil),
		},
		{
			Json:     `{}`,
			Expected: promocode.NewAgeRestriction(nil, nil, nil),
		},
		{
			Json: `{
				"lt": 30,
				"gt": 15
			}`,
			Expected: promocode.NewAgeRestriction(nil, ptr(15), ptr(30)),
		},
		{
			Json:     `{"gt": 30}`,
			Expected: promocode.NewAgeRestriction(nil, ptr(30), nil),
		},
		{
			Json:     `{"gt": 30`,
			Expected: nil,
		},
	}

	for i, tc := range testCases {
		var actual promocode.AgeRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		testParsing(t, i, tc, actual, err)
	}
}

func TestMeteoParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"is": "clear",
				"temp": {
					"gt": "15"
				}
			}`,
			Expected: promocode.NewMeteoRestriction("clear", nil, ptr(15), nil),
		},
		{
			Json: `{
				"is": "clear",
				"te`,
			Expected: nil,
		},
	}

	for i, tc := range testCases {
		var actual promocode.MeteoRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		testParsing(t, i, tc, actual, err)
	}
}

func TestDateParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"after": "2023-12-28",
				"before": "2023-12-30"
			}`,
			Expected: promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
		},
		{
			Json: `{
				"after": "2023-1",
			}`,
			Expected: nil,
		},
	}

	for i, tc := range testCases {
		var actual promocode.DateRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		testParsing(t, i, tc, actual, err)
	}
}

func TestAndParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `[
				{
					"@date": {
						"after": "2023-12-28",
						"before": "2023-12-30"
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
			]`,
			Expected: promocode.AndRestriction{
				Children: []promocode.Validator{
					promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
					promocode.NewMeteoRestriction("clear", nil, ptr(15), nil),
				},
			},
		},
		{
			Json: `[
				{
					"@and": [
						{
							"@date": {
								"after": "2023-12-28",
								"before": "2023-12-30"
							}
						}
					]
				}
			]`,
			Expected: promocode.AndRestriction{
				Children: []promocode.Validator{
					promocode.AndRestriction{
						Children: []promocode.Validator{
							promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		var actual promocode.AndRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		testParsing(t, i, tc, actual, err)
	}
}

func TestOrParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `[
				{
					"@date": {
						"after": "2023-12-28",
						"before": "2023-12-30"
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
			]`,
			Expected: promocode.OrRestriction{
				Children: []promocode.Validator{
					promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
					promocode.NewMeteoRestriction("clear", nil, ptr(15), nil),
				},
			},
		},
		{
			Json: `[
				{
					"@or": [
						{
							"@date": {
								"after": "2023-12-28",
								"before": "2023-12-30"
							}
						}
					]
				}
			]`,
			Expected: promocode.OrRestriction{
				Children: []promocode.Validator{
					promocode.OrRestriction{
						Children: []promocode.Validator{
							promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		var actual promocode.OrRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		testParsing(t, i, tc, actual, err)
	}
}
