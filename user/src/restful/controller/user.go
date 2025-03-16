package controller

import (
	"cloud-computing/users/restful/models/dto"
	"cloud-computing/users/restful/service"
	"cloud-computing/users/restful/validation"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.CreateUser(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// JoinOrganization handles adding a user to an organization
func JoinOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.JoinOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.JoinOrganization(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.JoinOrganization(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// LeaveOrganization handles removing a user from an organization
func LeaveOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.LeaveOrganizationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	if err := validation.LeaveOrganization(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	resp, err := service.LeaveOrganization(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
