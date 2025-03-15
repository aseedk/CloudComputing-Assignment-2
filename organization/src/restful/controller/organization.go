package controller

import (
	"cloud-computing/organization/organization/src/restful/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrganization handles organization creation
func CreateOrganization(c *gin.Context) {
	var req dto.CreateOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.CreateOrganization(req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	err := service.CreateOrganization(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, dto.CommonResponse{Code: 0, Message: "Organization created successfully", Data: nil})
}

// UpdateOrganization handles updating an organization
func UpdateOrganization(c *gin.Context) {
	var req dto.UpdateOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	var uri struct {
		OrganizationId string `uri:"organizationId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.UpdateOrganization(uri.OrganizationId, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	err := service.UpdateOrganization(uri.OrganizationId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, dto.CommonResponse{Code: 0, Message: "Organization updated successfully", Data: nil})
}

// DeleteOrganization handles deleting an organization
func DeleteOrganization(c *gin.Context) {
	var uri struct {
		OrganizationId string `uri:"organizationId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.DeleteOrganization(uri.OrganizationId); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	err := service.DeleteOrganization(uri.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, dto.CommonResponse{Code: 0, Message: "Organization deleted successfully", Data: nil})
}

// GetOrganization handles retrieving an organization
func GetOrganization(c *gin.Context) {
	var uri struct {
		OrganizationId string `uri:"organizationId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.GetOrganization(uri.OrganizationId); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	org, err := service.GetOrganization(uri.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, dto.CommonResponse{Code: 0, Message: "Success", Data: org})
}

// QueryOrganizations handles querying multiple organizations
func QueryOrganizations(c *gin.Context) {
	var req dto.QueryOrganizationReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.QueryOrganizations(req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	orgs, err := service.QueryOrganizations(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, dto.CommonResponse{Code: 0, Message: "Success", Data: orgs})
}
