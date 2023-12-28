package promocode

type Avantage struct {
	Percent int `json:"percent"`
}

type Promocode struct {
	Id           string         `json:"_id"`
	Name         string         `json:"name"`
	Avantage     Avantage       `json:"avantage"`
	Restrictions AndRestriction `json:"restrictions"`
}

func (p Promocode) Validate(args Arguments) (bool, error) {
	return p.Restrictions.Validate(args)
}
