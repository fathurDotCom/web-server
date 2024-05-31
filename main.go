package main

import (
	"web-server/config"
	"web-server/routes"
)

func main() {
	config.Connect()

	router := routes.SetupRouter()
	router.Run(":8080")
}
