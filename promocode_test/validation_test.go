package promocode_test

import (
	"testing"

	"github.com/thalkz/promo_code/promocode"
)

// This test argument will be used for all tests
var defaultTestArgument = promocode.Arguments{
	Age:         25,
	MeteoStatus: "clear",
	MeteoTemp:   15,
	Date:        parseDateOrPanic("2023-12-28"),
}

type testCase struct {
	Restriction promocode.Validator
	Expected    bool
}

func TestAgeRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.NewAgeRestriction(ptr(25), nil, nil),
			Expected:    true,
		},
		{
			Restriction: promocode.NewAgeRestriction(ptr(40), nil, nil),
			Expected:    false,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, nil, nil),
			Expected:    true,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, ptr(20), ptr(30)),
			Expected:    true,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, nil, ptr(40)),
			Expected:    true,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, nil, ptr(10)),
			Expected:    false,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, ptr(10), nil),
			Expected:    true,
		},
		{
			Restriction: promocode.NewAgeRestriction(nil, ptr(30), nil),
			Expected:    false,
		},
	}

	for i, tc := range testcases {
		valid, err := tc.Restriction.Validate(defaultTestArgument)
		if tc.Expected != valid {
			t.Errorf("validation failed for testcase #%v: expected %v (got %v, err: %v)", i, tc.Expected, valid, err)
		}
		// TODO Test if errors are thrown correctly
	}
}

func TestMeteoRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.NewMeteoRestriction("clear", nil, ptr(10), nil),
			Expected:    true,
		},
		{
			Restriction: promocode.NewMeteoRestriction("clear", nil, ptr(20), nil),
			Expected:    false,
		},
		{
			Restriction: promocode.NewMeteoRestriction("foggy", nil, ptr(10), nil),
			Expected:    false,
		},
	}

	for i, tc := range testcases {
		valid, err := tc.Restriction.Validate(defaultTestArgument)
		if tc.Expected != valid {
			t.Errorf("validation failed for testcase #%v: expected %v (got %v, err: %v)", i, tc.Expected, valid, err)
		}
		// TODO Test if errors are thrown correctly
	}
}

func TestDateRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.NewDateRestriction("2023-12-27", "2023-12-29"),
			Expected:    true,
		},
		{
			Restriction: promocode.NewDateRestriction("2023-12-28", "2023-12-28"),
			Expected:    true,
		},
		{
			Restriction: promocode.NewDateRestriction("", "2023-12-30"),
			Expected:    true,
		},
		{
			Restriction: promocode.NewDateRestriction("2023-12-30", ""),
			Expected:    false,
		},
		{
			Restriction: promocode.NewDateRestriction("", "2023-12-20"),
			Expected:    false,
		},
	}

	for i, tc := range testcases {
		valid, err := tc.Restriction.Validate(defaultTestArgument)
		if tc.Expected != valid {
			t.Errorf("validation failed for testcase #%v: expected %v (got %v, err: %v)", i, tc.Expected, valid, err)
		}
		// TODO Test if errors are thrown correctly
	}
}

func TestAndRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					validRestriction{},
				},
			},
			Expected: true,
		},
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					inalidRestriction{},
				},
			},
			Expected: false,
		},
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					inalidRestriction{},
				},
			},
			Expected: false,
		},
	}

	for i, tc := range testcases {
		valid, err := tc.Restriction.Validate(defaultTestArgument)
		if tc.Expected != valid {
			t.Errorf("validation failed for testcase #%v: expected %v (got %v, err: %v)", i, tc.Expected, valid, err)
		}
		// TODO Test if errors are thrown correctly
	}
}

func TestOrRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					validRestriction{},
				},
			},
			Expected: true,
		},
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					inalidRestriction{},
				},
			},
			Expected: true,
		},
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					inalidRestriction{},
					inalidRestriction{},
				},
			},
			Expected: false,
		},
	}

	for i, tc := range testcases {
		valid, err := tc.Restriction.Validate(defaultTestArgument)
		if tc.Expected != valid {
			t.Errorf("validation failed for testcase #%v: expected %v (got %v, err: %v)", i, tc.Expected, valid, err)
		}
		// TODO Test if errors are thrown correctly
	}
}
