package promocode

import (
	"errors"
	"fmt"
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
