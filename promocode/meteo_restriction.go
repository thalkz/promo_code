package promocode

import (
	"fmt"
)

type MeteoRestriction struct {
	Is   string             `json:"is"`
	Temp IntegerRestriction `json:"temp"`
}

func NewMeteoRestriction(is string, tempEq, tempGt, tempLt *int) MeteoRestriction {
	return MeteoRestriction{
		Is: is,
		Temp: IntegerRestriction{
			Eq: tempEq,
			Gt: tempGt,
			Lt: tempLt,
		},
	}
}

func (r MeteoRestriction) Validate(arg Arguments) (bool, error) {
	if r.Is != "" && r.Is != arg.MeteoStatus {
		return false, fmt.Errorf("invalid meteo status: expected %v (got %v)", r.Is, arg.MeteoStatus)
	}

	valid, err := r.Temp.InRange(arg.MeteoTemp)
	if err != nil || !valid {
		return false, fmt.Errorf("invalid temperature: %v", err)
	}

	return true, nil
}
