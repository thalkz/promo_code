package promocode

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateRestriction struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
}

// `before` and `after` need to be valid DateOnly time
func NewDateRestriction(after, before string) DateRestriction {
	afterTime, _ := time.Parse(time.DateOnly, after)
	beforeTime, _ := time.Parse(time.DateOnly, before)

	return DateRestriction{
		Before: beforeTime,
		After:  afterTime,
	}
}

func (r DateRestriction) Validate(arg Arguments) (bool, error) {
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

	d.After, err = time.Parse(time.DateOnly, result["after"])
	if err != nil {
		return fmt.Errorf("failed to parse after: %v", err)
	}

	d.Before, err = time.Parse(time.DateOnly, result["before"])
	if err != nil {
		return fmt.Errorf("failed to parse before: %v", err)
	}

	return err
}
