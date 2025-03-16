package route

import (
	"cloud-computing/users/restful/controller"
	"github.com/gin-gonic/gin"
)

// SetupUserRoute initializes user-related routes
func SetupUserRoute(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controller.CreateUser)             // Create a new user
		userGroup.POST("/join", controller.JoinOrganization)   // User joins an organization
		userGroup.POST("/leave", controller.LeaveOrganization) // User leaves an organization
	}
}
