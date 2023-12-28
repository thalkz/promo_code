package main

import "github.com/thalkz/promo_code/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
