package routersetup

import (
	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/routes"
)

// SetupRouters sets up the routes for the application
func SetupRouters(router *gin.Engine) {
    router.POST("/api/v1", routes.ShortenURL)
    router.GET("/api/v1/:shortID", routes.GetByShortID)
    router.DELETE("/api/v1/:shortID", routes.DeleteURL)
    router.PUT("/api/v1/:shortID", routes.EditURL)
    router.POST("/api/v1/addTag", routes.AddTag)
    router.GET("/api/v1/all-data", routes.GetAllData)
    router.GET("/api/v1/top-domains", routes.GetTopDomains)
}
