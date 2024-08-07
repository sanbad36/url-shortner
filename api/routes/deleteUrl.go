package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	err := database.Delete(shortID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the shortened Link"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shortened URL Deleted Successfully"})
}
