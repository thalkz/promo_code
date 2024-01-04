package promocode

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntegerRestriction is used both for `@age` and `@meteo[temp]`
type IntegerRestriction struct {
	Lt *int `json:"lt,omitempty"`
	Gt *int `json:"gt,omitempty"`
	Eq *int `json:"eq,omitempty"`
}

func (r IntegerRestriction) InRange(value int) (bool, error) {
	if r.Eq != nil && value != *r.Eq {
		return false, fmt.Errorf("should be equal to %v (got %v)", *r.Eq, value)
	}

	if r.Gt != nil && value < *r.Gt {
		return false, fmt.Errorf("should be greater than %v (got %v)", *r.Gt, value)
	}

	if r.Lt != nil && value > *r.Lt {
		return false, fmt.Errorf("should be less than %v (got %v)", *r.Lt, value)
	}

	return true, nil
}

func (r *IntegerRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]any
	err := json.Unmarshal(data, &result)

	eq, ok := result["eq"]
	if ok {
		ptr, err := parseIntOrString(eq)
		if err != nil {
			return err
		}
		r.Eq = ptr
	}

	gt, ok := result["gt"]
	if ok {
		ptr, err := parseIntOrString(gt)
		if err != nil {
			return err
		}
		r.Gt = ptr
	}

	lt, ok := result["lt"]
	if ok {
		ptr, err := parseIntOrString(lt)
		if err != nil {
			return err
		}
		r.Lt = ptr
	}

	return err
}

func parseIntOrString(input any) (*int, error) {
	switch input := input.(type) {
	case int:
		copy := int(input)
		return &copy, nil
	case float64:
		copy := int(input)
		return &copy, nil
	case string:
		v, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("cannot parse from unrecognized type: %T", input)
	}
}
