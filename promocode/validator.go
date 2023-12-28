package promocode

// Interface implemented by all restrictions to validate arguments
type Validator interface {
	Validate(arg Argument) (bool, error) // TODO Update arg type
}
