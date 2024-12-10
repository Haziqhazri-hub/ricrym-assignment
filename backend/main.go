package main

import (
	"database/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Run(":5432")
}