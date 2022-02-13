package middlewares

import "github.com/gin-gonic/gin"

func EditModuleMiddleware() gin.HandlerFunc {
	return basePermission("EDM")
}
