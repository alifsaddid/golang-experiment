package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateModule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   nil,
		"message": "Berhasil membuat module",
	})
}
