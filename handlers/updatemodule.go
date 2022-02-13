package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateModule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   nil,
		"message": "Berhasil mengupdate module",
	})
}
