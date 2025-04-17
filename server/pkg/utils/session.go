package utils

import (
	"log"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetSessionCookie(c *gin.Context, userID string, role string) {
	sessionID := "sess:user:" + userID
	value := userID + ":" + role

	err := config.RedisClient.Set(config.Ctx, sessionID, value, 24*time.Hour).Err()
	if err != nil {
		log.Println("failed to store session in Redis:", err)
	} else {
		log.Println("stored session:", sessionID, "→", value)
	}

	c.SetCookie("session_id", sessionID, 3600*24, "/", "", false, true)
}
