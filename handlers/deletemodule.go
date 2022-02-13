package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteModule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   nil,
		"message": "Berhasil menghapus module",
	})
}
