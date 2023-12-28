package promocode

import "time"

// Interface implemented by all restrictions to validate arguments
type Validator interface {
	Validate(args any) bool // TODO Update args type
}

type DateRestriction struct {
	After  time.Time
	Before time.Time
}

func (r DateRestriction) Validate(args any) bool {
	return true // TODO Implement me
}

type AgeExactRestriction struct {
	Eq int
}

func (r AgeExactRestriction) Validate(args any) bool {
	return true // TODO Implement me
}

type AgeRangeRestriction struct {
	Lt int
	Gt int
}

func (r AgeRangeRestriction) Validate(args any) bool {
	return true // TODO Implement me
}

type MeteoRestriction struct {
	Is   string
	Temp struct {
		Gt int
	}
}

func (r MeteoRestriction) Validate(args any) bool {
	return true // TODO Implement me
}

type AndRestriction struct {
	Children []Validator
}

func (r AndRestriction) Validate(args any) bool {
	return true // TODO Implement me
}

type OrRestriction struct {
	Children []Validator
}

func (r OrRestriction) Validate(args any) bool {
	return true // TODO Implement me
}
