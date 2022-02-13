package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModules(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   nil,
		"message": "Berhasil mendapatkan modules",
	})
}
