package main

import (
	"github.com/randytjioe/go-dts-07-randy/challenge_10/database"
	"github.com/randytjioe/go-dts-07-randy/challenge_10/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
