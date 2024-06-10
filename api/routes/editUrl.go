package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
	"github.com/sanbad36/url-shortner/api/models"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse JSON"})
		return
	}

	val, err := database.Get(shortID)
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "ShortID doesn't exists"})
		return
	}

	database.Set(shortID, body.URL, body.Expiry*3600*time.Second)
	c.JSON(http.StatusOK, gin.H{"message": "The content has been Updated"})
}
