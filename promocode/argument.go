package promocode

import "time"

type Argument struct {
	Age   int
	Date  time.Time
	Meteo struct {
		Town string
	}
}
