package controller

import (
	"cloud-computing/logging/restful/models/dto"
	"cloud-computing/logging/restful/service"
	"cloud-computing/logging/restful/validation"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLog handles log creation
func CreateLog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.CreateLogReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	// Validate request
	if err := validation.CreateLog(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	// Process request in service layer
	resp, err := service.CreateLog(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// QueryLogs handles querying logs
func QueryLogs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ApiContextTimeout)
	defer cancel()

	var req dto.QueryLogsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	// Validate request
	if err := validation.QueryLogs(ctx, &req); err != nil {
		c.JSON(http.StatusBadRequest, ParseError(err))
		return
	}

	// Fetch logs from service layer
	resp, err := service.QueryLogs(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ParseError(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
