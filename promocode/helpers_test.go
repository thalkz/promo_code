package promocode

import (
	"fmt"
	"time"
)

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

type validRestriction struct{}

func (r validRestriction) Validate(arg Argument) (bool, error) {
	return true, nil
}

type inalidRestriction struct{}

func (r inalidRestriction) Validate(arg Argument) (bool, error) {
	return false, fmt.Errorf("this restriction is always invalid")
}
