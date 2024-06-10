package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
	"github.com/sanbad36/url-shortner/api/models"
)

func EditURL(c *gin.Context){
	shortID := c.Param("shortID")
	var body models.Request
	
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error" : "Cannot Parse JSON",
		})
	}
	r := database.CreateClient(0)
	defer r.Close()

	// check if the shortId Exist in the DB or not

	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "ShortID doesn't exists",
		})
	}
	// update the content of the URL, expiry time with the shortID

	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Unable to update the shortend content",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "The content has been Updated",
	})
}