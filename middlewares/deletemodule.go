package middlewares

import "github.com/gin-gonic/gin"

func DeleteModuleMiddleware() gin.HandlerFunc {
	return basePermission("DLM")
}
