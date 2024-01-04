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

type validationTestCase struct {
	Restriction   promocode.Validator
	ExpectedValid bool
}

func testValidation(t *testing.T, caseId int, tc validationTestCase) {
	valid, err := tc.Restriction.Validate(defaultTestArgument)

	if tc.ExpectedValid {
		if !valid {
			t.Errorf("Testcase #%v: expected %v (got %v, err: %v)", caseId, tc.ExpectedValid, valid, err)
		}
	} else {
		if err == nil {
			t.Errorf("Testcase #%v: expected validation to throw an error, but didn't", caseId)
		}
	}
}

func TestAgeRestriction(t *testing.T) {
	var testcases = []validationTestCase{
		{
			Restriction:   promocode.NewAgeRestriction(ptr(25), nil, nil),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewAgeRestriction(ptr(40), nil, nil),
			ExpectedValid: false,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, nil, nil),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, ptr(20), ptr(30)),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, nil, ptr(40)),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, nil, ptr(10)),
			ExpectedValid: false,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, ptr(10), nil),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewAgeRestriction(nil, ptr(30), nil),
			ExpectedValid: false,
		},
	}

	for i, tc := range testcases {
		testValidation(t, i, tc)
	}
}

func TestMeteoRestriction(t *testing.T) {
	var testcases = []validationTestCase{
		{
			Restriction:   promocode.NewMeteoRestriction("clear", nil, ptr(10), nil),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewMeteoRestriction("clear", nil, ptr(20), nil),
			ExpectedValid: false,
		},
		{
			Restriction:   promocode.NewMeteoRestriction("foggy", nil, ptr(10), nil),
			ExpectedValid: false,
		},
	}

	for i, tc := range testcases {
		testValidation(t, i, tc)
	}
}

func TestDateRestriction(t *testing.T) {
	var testcases = []validationTestCase{
		{
			Restriction:   promocode.NewDateRestriction("2023-12-27", "2023-12-29"),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewDateRestriction("2023-12-28", "2023-12-28"),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewDateRestriction("", "2023-12-30"),
			ExpectedValid: true,
		},
		{
			Restriction:   promocode.NewDateRestriction("2023-12-30", ""),
			ExpectedValid: false,
		},
		{
			Restriction:   promocode.NewDateRestriction("", "2023-12-20"),
			ExpectedValid: false,
		},
	}

	for i, tc := range testcases {
		testValidation(t, i, tc)
	}
}

func TestAndRestriction(t *testing.T) {
	var testcases = []validationTestCase{
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					validRestriction{},
				},
			},
			ExpectedValid: true,
		},
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					inalidRestriction{},
				},
			},
			ExpectedValid: false,
		},
		{
			Restriction: promocode.AndRestriction{
				Children: []promocode.Validator{
					inalidRestriction{},
				},
			},
			ExpectedValid: false,
		},
	}

	for i, tc := range testcases {
		testValidation(t, i, tc)
	}
}

func TestOrRestriction(t *testing.T) {
	var testcases = []validationTestCase{
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					validRestriction{},
				},
			},
			ExpectedValid: true,
		},
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					validRestriction{},
					validRestriction{},
					inalidRestriction{},
				},
			},
			ExpectedValid: true,
		},
		{
			Restriction: promocode.OrRestriction{
				Children: []promocode.Validator{
					inalidRestriction{},
					inalidRestriction{},
				},
			},
			ExpectedValid: false,
		},
	}

	for i, tc := range testcases {
		testValidation(t, i, tc)
	}
}
