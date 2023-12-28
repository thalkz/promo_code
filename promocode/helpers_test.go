package promocode

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func ptr[T any](v T) *T {
	return &v
}

func assertSameJson(t *testing.T, testIndex int, expected any, actual any) {
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("TestCase #%v: failed to marshall expected: %v", testIndex, err)
	}
	expectedStr := string(expectedBytes)

	actualBytes, err := json.Marshal(actual)
	if err != nil {
		t.Errorf("TestCase #%v: failed to marshall actual: %v", testIndex, err)
	}
	actualStr := string(actualBytes)

	if expectedStr != actualStr {
		t.Errorf("TestCase #%v: expected %v (got %v)", testIndex, expectedStr, actualStr)
	}
}

func parseDateOrPanic(str string) time.Time {
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse date: %v", err))
	}
	return t
}

func createComplexTestRestriction() AndRestriction {
	var before, _ = time.Parse(time.DateOnly, "2019-01-01")
	var after, _ = time.Parse(time.DateOnly, "2020-06-30")
	return AndRestriction{
		Children: []Validator{
			DateRestriction{
				After:  before,
				Before: after,
			},
			OrRestriction{
				Children: []Validator{
					AgeRestriction{
						Eq: ptr(40),
					},
					AndRestriction{
						Children: []Validator{
							AgeRestriction{
								Lt: ptr(30),
								Gt: ptr(15),
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
			},
		},
	}
}

type validRestriction struct{}

func (r validRestriction) Validate(arg Arguments) (bool, error) {
	return true, nil
}

type inalidRestriction struct{}

func (r inalidRestriction) Validate(arg Arguments) (bool, error) {
	return false, fmt.Errorf("this restriction is always invalid")
}
