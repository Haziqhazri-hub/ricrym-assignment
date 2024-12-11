package routes

import (
	"net/http"

	"github.com/Haziqhazri-hub/ricrym-assignment/db"
	"github.com/Haziqhazri-hub/ricrym-assignment/model"
	"github.com/gin-gonic/gin"
)

// searchCharacters handles GET /search
func getUser(c *gin.Context) {
	// Get the search query from the query parameters
	query := c.DefaultQuery("query", "")

	// Retrieve the character search results
	ranks, err := model.SearchUser(db.DB, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if len(ranks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No characters found",
		})
		return
	}

	// Return the list of ranks and characters
	c.JSON(http.StatusOK, gin.H{
		"ranks": ranks,
	})
}
