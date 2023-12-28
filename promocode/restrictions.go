package promocode

import "time"

// Interface implemented by all restrictions to validate arguments
type Validator interface {
	Validate(arg Argument) bool // TODO Update arg type
}

type DateRestriction struct {
	After  time.Time
	Before time.Time
}

func (r DateRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}

type AgeExactRestriction struct {
	Eq int
}

func (r AgeExactRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}

type AgeRangeRestriction struct {
	Lt int
	Gt int
}

func (r AgeRangeRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}

type MeteoRestriction struct {
	Is   string
	Temp struct {
		Gt int
	}
}

func (r MeteoRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}

type AndRestriction struct {
	Children []Validator
}

func (r AndRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}

type OrRestriction struct {
	Children []Validator
}

func (r OrRestriction) Validate(arg Argument) bool {
	return true // TODO Implement me
}
