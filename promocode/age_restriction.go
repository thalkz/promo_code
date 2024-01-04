package promocode

import (
	"encoding/json"
	"fmt"
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
	valid, err := r.InRange(arg.Age)
	if err != nil || !valid {
		return false, fmt.Errorf("invalid age: %w", err)
	}
	return true, nil
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
