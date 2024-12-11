package main

import (
	"github.com/Haziqhazri-hub/ricrym-assignment/db"
	"github.com/Haziqhazri-hub/ricrym-assignment/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
