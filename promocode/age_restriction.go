package promocode

import (
	"encoding/json"
)

type AgeRestriction struct {
	IntegerRestriction
}

func NewAgeRestriction(eq, gt, lt *int) AgeRestriction {
	return AgeRestriction{
		IntegerRestriction: IntegerRestriction{
			Eq: eq,
			Gt: gt,
			Lt: lt,
		},
	}
}

func (r AgeRestriction) Validate(arg Arguments) (bool, error) {
	return r.InRange(arg.Age)
}

func (r *AgeRestriction) UnmarshalJSON(data []byte) error {
	var content IntegerRestriction
	err := json.Unmarshal(data, &content)
	if err != nil {
		return err
	}
	r.IntegerRestriction = content
	return nil
}
