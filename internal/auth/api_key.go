package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	apiKey := os.Getenv("API_KEY_AUTH")
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-KEY")
		if key != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
