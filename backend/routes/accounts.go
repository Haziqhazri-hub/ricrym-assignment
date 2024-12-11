package routes

import (
	"github.com/Haziqhazri-hub/ricrym-assignment/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAccounts(c *gin.Context) {
	accounts, err := model.GetAllAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)

}