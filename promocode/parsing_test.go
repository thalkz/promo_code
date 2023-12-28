package promocode

import (
	"encoding/json"
	"math"
	"testing"
)

type parsingTestCase struct {
	Json       string
	Expected   Validator
	ShouldFail bool
}

func TestAgeRangeParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"lt": 30,
				"gt": 15
			}`,
			Expected: AgeRangeRestriction{
				Lt: 30,
				Gt: 15,
			},
			ShouldFail: false,
		},
		{
			Json: `{}`,
			Expected: AgeRangeRestriction{
				Gt: 0,
				Lt: math.MaxInt,
			},
			ShouldFail: false,
		},
		{
			Json: `{"gt": 30}`,
			Expected: AgeRangeRestriction{
				Gt: 30,
				Lt: math.MaxInt,
			},
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual AgeRangeRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if tc.ShouldFail && err == nil {
			t.Errorf("TestCase #%v: failed: no error was thrown", i)
		}
		if !tc.ShouldFail && err != nil {
			t.Errorf("TestCase #%v: failed with an error: %v", i, err)
		}
		if actual != tc.Expected {
			t.Errorf("TestCase #%v: failed to parse input: expected %v (got %v)", i, tc.Expected, actual)
		}
	}
}

func TestAgeExactParsing(t *testing.T) {
	var testCases = []parsingTestCase{
		{
			Json: `{
				"eq": 30
			}`,
			Expected: AgeExactRestriction{
				Eq: 30,
			},
			ShouldFail: false,
		},
		{
			Json: `{}`,
			Expected: AgeExactRestriction{
				Eq: 0,
			},
			ShouldFail: false,
		},
	}

	for i, tc := range testCases {
		var actual AgeExactRestriction
		err := json.Unmarshal([]byte(tc.Json), &actual)
		if tc.ShouldFail && err == nil {
			t.Errorf("TestCase #%v: failed: no error was thrown", i)
		}
		if !tc.ShouldFail && err != nil {
			t.Errorf("TestCase #%v: failed with an error: %v", i, err)
		}
		if actual != tc.Expected {
			t.Errorf("TestCase #%v: failed to parse input: expected %v (got %v)", i, tc.Expected, actual)
		}
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
		if tc.ShouldFail && err == nil {
			t.Errorf("TestCase #%v: failed: no error was thrown", i)
		}
		if !tc.ShouldFail && err != nil {
			t.Errorf("TestCase #%v: failed with an error: %v", i, err)
		}
		if actual != tc.Expected {
			t.Errorf("TestCase #%v: failed to parse input: expected %v (got %v)", i, tc.Expected, actual)
		}
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
		if tc.ShouldFail && err == nil {
			t.Errorf("TestCase #%v: failed: no error was thrown", i)
		}
		if !tc.ShouldFail && err != nil {
			t.Errorf("TestCase #%v: failed with an error: %v", i, err)
		}
		if actual != tc.Expected {
			t.Errorf("TestCase #%v: failed to parse input: expected %v (got %v)", i, tc.Expected, actual)
		}
	}
}
