package middleware

import (
	"net/http"
	"strings"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: missing session"})
			c.Abort()
			return
		}

		val, err := config.RedisClient.Get(config.Ctx, sessionID).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: invalid session"})
			c.Abort()
			return
		}

		parts := strings.Split(val, ":")
		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: session format invalid"})
			c.Abort()
			return
		}

		c.Set("userID", parts[0])
		c.Set("role", parts[1])
		c.Next()

	}
}
