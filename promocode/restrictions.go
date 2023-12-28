package promocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Interface implemented by all restrictions to validate arguments
type Validator interface {
	Validate(arg Argument) (bool, error) // TODO Update arg type
}

type AnyRestriction struct {
	*DateRestriction  `json:"@date"`
	*MeteoRestriction `json:"@meteo"`
	*AgeRestriction   `json:"@age"`
	*AndRestriction   `json:"@and"`
	*OrRestriction    `json:"@or"`
}

func (r AnyRestriction) Find() (Validator, error) {
	if r.AgeRestriction != nil {
		return *r.AgeRestriction, nil
	} else if r.MeteoRestriction != nil {
		return *r.MeteoRestriction, nil
	} else if r.DateRestriction != nil {
		return *r.DateRestriction, nil
	} else if r.AndRestriction != nil {
		return *r.AndRestriction, nil
	} else if r.OrRestriction != nil {
		return *r.OrRestriction, nil
	}
	return nil, fmt.Errorf("failed to find validator in any restriction")
}

type DateRestriction struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
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

// type AgeRestriction struct {
// 	Eq int
// }

// func (r AgeRestriction) Validate(arg Argument) (bool, error) {
// 	if r.Eq != arg.Age {
// 		return false, fmt.Errorf("invalid age: expected %v (got %v)", r.Eq, arg.Age)
// 	}
// 	return true, nil
// }

type AgeRestriction struct {
	Lt *int
	Gt *int
	Eq *int
}

func (r AgeRestriction) Validate(arg Argument) (bool, error) {
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

func (d *AndRestriction) UnmarshalJSON(data []byte) error {
	var result []AnyRestriction
	err := json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("failed to parse json: %v", err)
	}

	d.Children = make([]Validator, 0)
	for _, anyRestriction := range result {
		item, err := anyRestriction.Find()
		if err != nil {
			return fmt.Errorf("failed to parse any restriction: %v", err)
		}
		d.Children = append(d.Children, item)
	}

	return err
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

func (d *OrRestriction) UnmarshalJSON(data []byte) error {
	var result []AnyRestriction
	err := json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("failed to parse json: %v", err)
	}

	d.Children = make([]Validator, 0)
	for _, anyRestriction := range result {
		item, err := anyRestriction.Find()
		if err != nil {
			return fmt.Errorf("failed to parse any restriction: %v", err)
		}
		d.Children = append(d.Children, item)
	}

	return err
}
