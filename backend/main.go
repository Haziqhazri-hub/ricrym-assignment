package main

import (
	"os"

	"github.com/Haziqhazri-hub/ricrym-assignment/db"
	"github.com/Haziqhazri-hub/ricrym-assignment/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Use(cors.Default())

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
