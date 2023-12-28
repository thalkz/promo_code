package promocode

import (
	"encoding/json"
	"errors"
	"fmt"
)

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
