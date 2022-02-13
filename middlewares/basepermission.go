package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func basePermission(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Permissions"] == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "No permission",
				"message": "Tidak ada permission",
			})

			return
		}

		permissions := strings.Split(c.Request.Header["Permissions"][0], ",")
		if !contains(permissions, p) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "No permission",
				"message": "Tidak ada permission",
			})

			return
		}

		c.Next()
	}
}

func contains(sl []string, s string) bool {
	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}
