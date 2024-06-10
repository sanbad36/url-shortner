package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
)

func GetByShortID(c *gin.Context) {
	shortID := c.Param("shortID")

	val, err := database.Get(shortID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for given ShortID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": val})
}
