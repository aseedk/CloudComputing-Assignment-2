package route

import (
	"cloud-computing/organization/organization/src/restful/controller"
	"github.com/gin-gonic/gin"
)

// SetupOrganizationRoute initializes organization-related routes
func SetupOrganizationRoute(router *gin.Engine) {
	organizationGroup := router.Group("/organizations")
	{
		organizationGroup.POST("/", controller.CreateOrganization)
		organizationGroup.PUT("/:organizationId", controller.UpdateOrganization)
		organizationGroup.DELETE("/:organizationId", controller.DeleteOrganization)
		organizationGroup.GET("/:organizationId", controller.GetOrganization)
		organizationGroup.GET("/", controller.QueryOrganizations)
	}
}
