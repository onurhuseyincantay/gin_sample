package Handlers

import "github.com/gin-gonic/gin"

func ErrorHandler(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}
