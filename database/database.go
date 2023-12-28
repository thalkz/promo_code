package database

import "github.com/thalkz/promo_code/promocode"

// Simple in memory database
var Instance = map[string]*promocode.Promocode{}

func Reset() {
	for k := range Instance {
		delete(Instance, k)
	}
}
