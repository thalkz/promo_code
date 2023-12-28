package promocode

import (
	"encoding/json"
	"testing"
)

type parsingTestCase struct {
	Json       string
	Expected   Validator
	ShouldFail bool
}

func TestAgeParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"eq": 30
			}`,
			Expected: AgeRestriction{
				Eq: ptr(30),
			},
			ShouldFail: false,
		},
		{
			Json:       `{}`,
			Expected:   AgeRestriction{},
			ShouldFail: false,
		},
		{
			Json: `{
				"lt": 30,
				"gt": 15
			}`,
			Expected: AgeRestriction{
				Lt: ptr(30),
				Gt: ptr(15),
			},
			ShouldFail: false,
		},
		{
			Json: `{"gt": 30}`,
			Expected: AgeRestriction{
				Gt: ptr(30),
			},
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual AgeRestriction
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
			Expected: MeteoRestriction{
				Is: "clear",
				Temp: struct{ Gt int }{
					Gt: 15,
				},
			},
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual MeteoRestriction
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
			Expected: DateRestriction{
				Before: parseDateOrPanic("2023-12-30"),
				After:  parseDateOrPanic("2023-12-28"),
			},
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual DateRestriction
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
			Expected: AndRestriction{
				Children: []Validator{
					DateRestriction{
						Before: parseDateOrPanic("2023-12-30"),
						After:  parseDateOrPanic("2023-12-28"),
					},
					MeteoRestriction{
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
			Expected: AndRestriction{
				Children: []Validator{
					AndRestriction{
						Children: []Validator{
							DateRestriction{
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
		var actual AndRestriction
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
			Expected: OrRestriction{
				Children: []Validator{
					DateRestriction{
						Before: parseDateOrPanic("2023-12-30"),
						After:  parseDateOrPanic("2023-12-28"),
					},
					MeteoRestriction{
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
			Expected: OrRestriction{
				Children: []Validator{
					OrRestriction{
						Children: []Validator{
							DateRestriction{
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
		var actual OrRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if err != nil {
			t.Errorf("TestCase #%v: failed to parse json: %v", i, err)
		}
		assertSameJson(t, i, tc.Expected, actual)
	}
}
