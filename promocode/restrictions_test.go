package promocode

import (
	"testing"
	"time"
)

func createTestRestriction() AndRestriction {
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
					AgeExactRestriction{
						Eq: 40,
					},
					AndRestriction{
						Children: []Validator{
							AgeRangeRestriction{
								Lt: 30,
								Gt: 15,
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

func TestRestriction(t *testing.T) {
	restriction := createTestRestriction()
	if restriction.Validate(nil) {
		t.Errorf("validate should return true")
	}
}
