package main

import (
	"log"
	"myapp/database"
	"myapp/router"
)

func main() {
	database.ConnectDB()
	database.Migrate()

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
