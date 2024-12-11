package routes

import (
	"net/http"
	"github.com/Haziqhazri-hub/ricrym-assignment/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func getPagination(c *gin.Context) {
	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize := 6

	ranks, totalPages, err := model.GetPaginatedRank(page, pageSize)
	if err != nil {
		log.Printf("Error fetching paginated ranks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"totalPages": totalPages,
		"data":       ranks,
	})
}
