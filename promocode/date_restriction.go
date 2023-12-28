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

	d.After, _ = time.Parse(time.DateOnly, result["after"])
	d.Before, _ = time.Parse(time.DateOnly, result["before"])

	return err
}
