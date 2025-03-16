package route

import (
	"cloud-computing/logging/restful/controller"
	"github.com/gin-gonic/gin"
)

// SetupLoggingRoutes initializes logging-related routes
func SetupLoggingRoutes(router *gin.Engine) {
	loggingGroup := router.Group("/logging")
	{
		loggingGroup.POST("/", controller.CreateLog)
		loggingGroup.GET("/", controller.QueryLogs)
	}
}
