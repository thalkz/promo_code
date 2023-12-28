package database

import "github.com/thalkz/promo_code/promocode"

// Simple in memory database
var Instance = map[string]*promocode.AndRestriction{}
