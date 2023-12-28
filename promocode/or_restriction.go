package promocode

import (
	"encoding/json"
	"errors"
	"fmt"
)

type OrRestriction struct {
	Children []Validator
}

func (r OrRestriction) Validate(arg Arguments) (bool, error) {
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
