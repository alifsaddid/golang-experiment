package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CreateModuleMiddleware() gin.HandlerFunc {
	return basePermission("CRM")
}
