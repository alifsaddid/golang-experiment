package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	// Pura-puranya ambil list permission dari database
	c.Request.Header.Add("Permissions", "CRM,EDM")
	c.Next()
}
