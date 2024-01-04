package promocode_test

import (
	"fmt"
	"time"

	"github.com/thalkz/promo_code/promocode"
)

type validRestriction struct{}

func (r validRestriction) Validate(arg promocode.Arguments) (bool, error) {
	return true, nil
}

type inalidRestriction struct{}

func (r inalidRestriction) Validate(arg promocode.Arguments) (bool, error) {
	return false, fmt.Errorf("this restriction is always invalid")
}

func ptr[T any](v T) *T {
	return &v
}

func parseDateOrPanic(str string) time.Time {
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse date: %v", err))
	}
	return t
}
