package main

import (
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.InitDB()
	defer db.DB.Close()

	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			db.ExpireSeats()
		}
	}()

	routes.AllRoutes(server)

	server.Run(":8080")

}
