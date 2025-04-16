package utils

import (
	"os"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetSessionCookie(c *gin.Context, userID string) {
	if config.RedisClient == nil || os.Getenv("TEST_MODE") == "true" {
		return
	}

	sessionID := "sess:user:" + userID

	config.RedisClient.Set(config.Ctx, sessionID, userID, 24*time.Hour)
	c.SetCookie("session_id", sessionID, 3600*24, "/", "", false, true)
}
