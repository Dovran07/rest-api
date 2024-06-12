package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

type statusRespone struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": message,
	})
	return
}

func newSuccessResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
	return
}
