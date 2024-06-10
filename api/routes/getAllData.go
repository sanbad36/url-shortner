// routes/data.go

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
	"github.com/sanbad36/url-shortner/api/database/store"
)

func GetAllData(c *gin.Context) {
	var allData map[string]string
	var err error

	if database.IsRedisAvailable() {
		allData, err = getAllDataFromRedis()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from Redis"})
			return
		}
	} else {
		allData = store.GetAllDataFromInMemoryStore()
	}

	// Filter out deleted entries from allData
	filteredData := make(map[string]string)
	for key, value := range allData {
		_, exists := store.Get(key)
		if exists {
			filteredData[key] = value
		}
	}

	c.JSON(http.StatusOK, filteredData)
}

func getAllDataFromRedis() (map[string]string, error) {
	keys, err := database.Keys("*")
	if err != nil {
		return nil, err
	}

	allData := make(map[string]string)
	for _, key := range keys {
		val, err := database.Get(key)
		if err != nil {
			continue
		}
		allData[key] = val
	}

	return allData, nil
}
