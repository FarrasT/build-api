package main

import (
	"build-api/database"
	"build-api/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
