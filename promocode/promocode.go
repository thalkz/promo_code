package promocode

type Avantage struct {
	Percent int
}

type Promocode struct {
	Id           string
	Name         string
	Avantage     Avantage
	Restrictions AndRestriction
}

func (p Promocode) Validate(args Arguments) (bool, error) {
	return p.Restrictions.Validate(args)
}
