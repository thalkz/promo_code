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
			Expected: promocode.AgeRestriction{
				Eq: ptr(30),
			},
			ShouldFail: false,
		},
		{
			Json:       `{}`,
			Expected:   promocode.AgeRestriction{},
			ShouldFail: false,
		},
		{
			Json: `{
				"lt": 30,
				"gt": 15
			}`,
			Expected: promocode.AgeRestriction{
				Lt: ptr(30),
				Gt: ptr(15),
			},
			ShouldFail: false,
		},
		{
			Json: `{"gt": 30}`,
			Expected: promocode.AgeRestriction{
				Gt: ptr(30),
			},
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
			Expected: promocode.MeteoRestriction{
				Is: "clear",
				Temp: struct{ Gt int }{
					Gt: 15,
				},
			},
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
			Expected: promocode.DateRestriction{
				Before: parseDateOrPanic("2023-12-30"),
				After:  parseDateOrPanic("2023-12-28"),
			},
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
					promocode.DateRestriction{
						Before: parseDateOrPanic("2023-12-30"),
						After:  parseDateOrPanic("2023-12-28"),
					},
					promocode.MeteoRestriction{
						Is: "clear",
						Temp: struct{ Gt int }{
							Gt: 15,
						},
					},
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
							promocode.DateRestriction{
								After:  parseDateOrPanic("2023-12-28"),
								Before: parseDateOrPanic("2023-12-30"),
							},
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
					promocode.DateRestriction{
						Before: parseDateOrPanic("2023-12-30"),
						After:  parseDateOrPanic("2023-12-28"),
					},
					promocode.MeteoRestriction{
						Is: "clear",
						Temp: struct{ Gt int }{
							Gt: 15,
						},
					},
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
							promocode.DateRestriction{
								After:  parseDateOrPanic("2023-12-28"),
								Before: parseDateOrPanic("2023-12-30"),
							},
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
