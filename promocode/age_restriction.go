package promocode

import (
	"encoding/json"
	"fmt"
)

type AgeRestriction struct {
	Lt *int
	Gt *int
	Eq *int
}

func (r AgeRestriction) Validate(arg Arguments) (bool, error) {
	if r.Eq != nil && arg.Age != *r.Eq {
		return false, fmt.Errorf("invalid age: should be equal to %v (got %v)", *r.Eq, arg.Age)
	}

	if r.Gt != nil && arg.Age < *r.Gt {
		return false, fmt.Errorf("invalid age: should be greater than %v (got %v)", *r.Gt, arg.Age)
	}
	if r.Lt != nil && arg.Age > *r.Lt {
		return false, fmt.Errorf("invalid age: should be less than %v (got %v)", *r.Lt, arg.Age)
	}

	return true, nil
}

func (r *AgeRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]int
	err := json.Unmarshal(data, &result)

	eq, ok := result["eq"]
	if ok {
		r.Eq = &eq
	}

	gt, ok := result["gt"]
	if ok {
		r.Gt = &gt
	}

	lt, ok := result["lt"]
	if ok {
		r.Lt = &lt
	}

	return err
}
