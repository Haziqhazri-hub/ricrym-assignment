package routes

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine) {
	server.GET("/accounts", getAccounts)
	server.GET("/characters", getCharacters)
	server.GET("/pagination/:page", getPagination)
	server.GET("/search/:username", getUser)
}
