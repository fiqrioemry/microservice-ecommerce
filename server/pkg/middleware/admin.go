package middleware

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.MustGetRole(c)

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden - Admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
