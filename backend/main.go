package main

import (
	"github.com/Haziqhazri-hub/ricrym-assignment/db"
	"github.com/Haziqhazri-hub/ricrym-assignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}