package main

import (
	"github.com/joho/godotenv"
	"github.com/thalkz/promo_code/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	r := router.SetupRouter()
	r.Run(":8080")
}
