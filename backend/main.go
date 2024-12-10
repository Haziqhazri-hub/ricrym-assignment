package main

import (
	"database/db"
	"database/model"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	model.GenerateFakeData()

	server.Run(":5432")
}