package test_helper

import (
	"encoding/json"
	"testing"
)

func AssertSameJson(t *testing.T, testIndex int, expected any, actual any) {
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
