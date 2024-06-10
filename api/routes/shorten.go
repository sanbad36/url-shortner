package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sanbad36/url-shortner/api/database"
	"github.com/sanbad36/url-shortner/api/models"
	"github.com/sanbad36/url-shortner/api/utils"
)

func ShortenURL(c *gin.Context) {
	var body models.Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}

	val, err := database.Get(c.ClientIP())
	if err != nil {
		database.Set(c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second)
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := database.TTL(c.ClientIP())
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}

	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	if !utils.IsDifferentDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "You can't hack this System :)",
		})
		return
	}
	body.URL = utils.EnsureHttpPrefix(body.URL)

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	val, _ = database.Get(id)
	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "URL Custom Short Already Exists",
		})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}
	database.Set(id, body.URL, body.Expiry*3600*time.Second)

	resp := models.Response{
		Expiry:         body.Expiry,
		XRateRemaining: 10,
		XRateLimitReset: 30,
		URL:            body.URL,
		CustomShort:    "",
	}
	database.Decr(c.ClientIP())
	val, _ = database.Get(c.ClientIP())
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := database.TTL(c.ClientIP())
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id
	c.JSON(http.StatusOK, resp)
}
