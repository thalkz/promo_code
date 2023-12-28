package promocode

import "fmt"

type AnyRestriction struct {
	*DateRestriction  `json:"@date"`
	*MeteoRestriction `json:"@meteo"`
	*AgeRestriction   `json:"@age"`
	*AndRestriction   `json:"@and"`
	*OrRestriction    `json:"@or"`
}

func (r AnyRestriction) Find() (Validator, error) {
	if r.AgeRestriction != nil {
		return *r.AgeRestriction, nil
	} else if r.MeteoRestriction != nil {
		return *r.MeteoRestriction, nil
	} else if r.DateRestriction != nil {
		return *r.DateRestriction, nil
	} else if r.AndRestriction != nil {
		return *r.AndRestriction, nil
	} else if r.OrRestriction != nil {
		return *r.OrRestriction, nil
	}
	return nil, fmt.Errorf("failed to find validator in any restriction")
}
