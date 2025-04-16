package middleware

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/pkg/config"

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

		userID, err := config.RedisClient.Get(config.Ctx, sessionID).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: invalid session"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
