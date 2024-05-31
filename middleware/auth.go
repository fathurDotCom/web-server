package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "your-secret-token" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
