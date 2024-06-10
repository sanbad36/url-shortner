package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest
	if err := c.ShouldBind(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}
	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	val, err := database.Get(shortID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for the given ShortID"})
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		data = make(map[string]interface{})
		data["data"] = val
	}

	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag Already Exists"})
			return
		}
	}
	tags = append(tags, tag)
	data["tags"] = tags

	updateData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Marshal updated data"})
		return
	}
	database.Set(shortID, string(updateData), 0)
	c.JSON(http.StatusOK, data)
}
