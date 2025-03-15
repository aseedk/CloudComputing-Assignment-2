package controller

import (
	"cloud-computing/organization/organization/src/restful/models/dto"
	"cloud-computing/organization/organization/src/restful/service"
	"cloud-computing/organization/organization/src/restful/validation"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrganization handles organization creation
func CreateOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.CreateOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.CreateOrganization(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.CreateOrganization(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateOrganization handles updating an organization
func UpdateOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.UpdateOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.UpdateOrganization(ctx, req.OrganizationId); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.UpdateOrganization(ctx, req.OrganizationId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteOrganization handles deleting an organization
func DeleteOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var uri struct {
		OrganizationId string `uri:"organizationId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.DeleteOrganization(ctx, uri.OrganizationId); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.DeleteOrganization(ctx, uri.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetOrganization handles retrieving an organization
func GetOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var uri struct {
		OrganizationId string `uri:"organizationId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.GetOrganization(ctx, uri.OrganizationId); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.GetOrganization(ctx, uri.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// QueryOrganizations handles querying multiple organizations
func QueryOrganizations(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.QueryOrganizationReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.QueryOrganizations(ctx, &req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.QueryOrganizations(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
