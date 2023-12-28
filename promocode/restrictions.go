package promocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Interface implemented by all restrictions to validate arguments
type Validator interface {
	Validate(arg Argument) (bool, error) // TODO Update arg type
}

type DateRestriction struct {
	After  time.Time
	Before time.Time
}

func (r DateRestriction) Validate(arg Argument) (bool, error) {
	if arg.Date.IsZero() {
		return false, fmt.Errorf("missing date in argument")
	}

	if !r.After.IsZero() && arg.Date.Before(r.After) {
		return false, fmt.Errorf("promocode is not valid yet")
	}

	if !r.Before.IsZero() && arg.Date.After(r.Before) {
		return false, fmt.Errorf("promocode is not valid anymore")
	}

	return true, nil
}

func (d *DateRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("failed to parse date restriction: %v", err)
	}

	d.After, _ = time.Parse(time.DateOnly, result["after"])
	d.Before, _ = time.Parse(time.DateOnly, result["before"])

	return err
}

type AgeExactRestriction struct {
	Eq int
}

func (r AgeExactRestriction) Validate(arg Argument) (bool, error) {
	if r.Eq != arg.Age {
		return false, fmt.Errorf("invalid age: expected %v (got %v)", r.Eq, arg.Age)
	}
	return true, nil
}

type AgeRangeRestriction struct {
	Lt int
	Gt int
}

func (r AgeRangeRestriction) Validate(arg Argument) (bool, error) {
	if arg.Age < r.Gt {
		return false, fmt.Errorf("invalid age: should be greater than %v (got %v)", r.Gt, arg.Age)
	}
	if arg.Age > r.Lt {
		return false, fmt.Errorf("invalid age: should be less than %v (got %v)", r.Lt, arg.Age)
	}

	return true, nil
}

func (d *AgeRangeRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]int
	err := json.Unmarshal(data, &result)

	d.Gt = result["gt"]
	d.Lt = result["lt"]

	// Default to MaxInt for Lt
	// TODO Make sure this is correct behavior
	if d.Lt == 0 {
		d.Lt = math.MaxInt
	}

	return err
}

type MeteoRestriction struct {
	Is   string
	Temp struct {
		Gt int
	}
}

func (r MeteoRestriction) Validate(arg Argument) (bool, error) {
	if r.Is != "" && r.Is != arg.MeteoStatus {
		return false, fmt.Errorf("invalid meteo status: expected %v (got %v)", r.Is, arg.MeteoStatus)
	}

	if arg.MeteoTemp < r.Temp.Gt {
		return false, fmt.Errorf("invalid temperature: should be greater than %v (got %v)", r.Temp.Gt, arg.Age)
	}

	return true, nil
}

func (d *MeteoRestriction) UnmarshalJSON(data []byte) error {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("failed to parse json: %v", err)
	}

	d.Is = result["is"].(string)
	temp := result["temp"].(map[string]any)
	gtStr := temp["gt"].(string)
	d.Temp.Gt, err = strconv.Atoi(gtStr)

	if err != nil {
		return fmt.Errorf("failed to parse gt: %v", err)
	}

	return err
}

type AndRestriction struct {
	Children []Validator
}

func (r AndRestriction) Validate(arg Argument) (bool, error) {
	var result = true
	var errs = make([]error, 0)
	for _, child := range r.Children {
		valid, err := child.Validate(arg)

		// If any condition is false, the validation fails
		if !valid {
			result = false
		}

		// Gather all errors to return them all
		if err != nil {
			errs = append(errs, err)
		}
	}
	allErrors := errors.Join(errs...)
	return result, fmt.Errorf("failed AND condition: %v", allErrors)
}

type OrRestriction struct {
	Children []Validator
}

func (r OrRestriction) Validate(arg Argument) (bool, error) {
	var result = false
	var errs = make([]error, 0)

	for _, child := range r.Children {
		valid, err := child.Validate(arg)

		// If any condition is true, the validation succeeds
		if valid {
			result = true
		}

		// Gather all errors to return them all
		if err != nil {
			errs = append(errs, err)
		}
	}

	allErrors := errors.Join(errs...)
	return result, fmt.Errorf("failed OR condition: %v", allErrors)
}
