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
			Restriction: promocode.AgeRestriction{
				Eq: ptr(25),
			},
			Expected: true,
		},
		{
			Restriction: promocode.AgeRestriction{
				Eq: ptr(40),
			},
			Expected: false,
		},
		{
			Restriction: promocode.AgeRestriction{},
			Expected:    true,
		},
		{
			Restriction: promocode.AgeRestriction{
				Gt: ptr(20),
				Lt: ptr(30),
			},
			Expected: true,
		},
		{
			Restriction: promocode.AgeRestriction{
				Lt: ptr(40),
			},
			Expected: true,
		},
		{
			Restriction: promocode.AgeRestriction{
				Lt: ptr(10),
			},
			Expected: false,
		},
		{
			Restriction: promocode.AgeRestriction{
				Gt: ptr(10),
			},
			Expected: true,
		},
		{
			Restriction: promocode.AgeRestriction{
				Gt: ptr(30),
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

func TestMeteoRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.MeteoRestriction{
				Is: "clear",
				Temp: struct{ Gt int }{
					Gt: 10,
				},
			},
			Expected: true,
		},
		{
			Restriction: promocode.MeteoRestriction{
				Is: "clear",
				Temp: struct{ Gt int }{
					Gt: 20,
				},
			},
			Expected: false,
		},
		{
			Restriction: promocode.MeteoRestriction{
				Is: "foggy",
				Temp: struct{ Gt int }{
					Gt: 10,
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

func TestDateRestriction(t *testing.T) {
	var testcases = []testCase{
		{
			Restriction: promocode.DateRestriction{
				After:  parseDateOrPanic("2023-12-27"),
				Before: parseDateOrPanic("2023-12-29"),
			},
			Expected: true,
		},
		{
			Restriction: promocode.DateRestriction{
				After:  parseDateOrPanic("2023-12-28"),
				Before: parseDateOrPanic("2023-12-28"),
			},
			Expected: true,
		},
		{
			Restriction: promocode.DateRestriction{
				Before: parseDateOrPanic("2023-12-30"),
			},
			Expected: true,
		},
		{
			Restriction: promocode.DateRestriction{
				After: parseDateOrPanic("2023-12-20"),
			},
			Expected: true,
		},
		{
			Restriction: promocode.DateRestriction{
				After: parseDateOrPanic("2023-12-30"),
			},
			Expected: false,
		},
		{
			Restriction: promocode.DateRestriction{
				Before: parseDateOrPanic("2023-12-20"),
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
