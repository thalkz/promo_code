package promocode_test

import (
	"encoding/json"
	"testing"

	"github.com/thalkz/promo_code/promocode"
)

type parsingTestCase struct {
	Json       string
	Expected   promocode.Validator
	ShouldFail bool
}

func TestAgeParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"eq": 30
			}`,
			Expected:   promocode.NewAgeRestriction(ptr(30), nil, nil),
			ShouldFail: false,
		},
		{
			Json:       `{}`,
			Expected:   promocode.NewAgeRestriction(nil, nil, nil),
			ShouldFail: false,
		},
		{
			Json: `{
				"lt": 30,
				"gt": 15
			}`,
			Expected:   promocode.NewAgeRestriction(nil, ptr(15), ptr(30)),
			ShouldFail: false,
		},
		{
			Json:       `{"gt": 30}`,
			Expected:   promocode.NewAgeRestriction(nil, ptr(30), nil),
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual promocode.AgeRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
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
			Expected:   promocode.NewMeteoRestriction("clear", nil, ptr(15), nil),
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual promocode.MeteoRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
	}
}

func TestDateParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"after": "2023-12-28",
				"before": "2023-12-30"
			}`,
			Expected:   promocode.NewDateRestriction("2023-12-28", "2023-12-30"),
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual promocode.DateRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
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
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
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
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
	}
}
