package routes

import (
	"github.com/Haziqhazri-hub/ricrym-assignment/model"
	"github.com/gin-gonic/gin"
)

func getAllRanksHandler(c *gin.Context) {
	ranks, err := model.GetAllRanks()
	if err != nil {
		c.JSON(500, gin.H{"message": "Error fetching ranks", "error": err.Error()})
		return
	}

	c.JSON(200, ranks)
}
