package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MustGetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		panic("userID not found in context")
	}
	idStr, ok := userID.(string)
	if !ok {
		panic("userID in context is not a string")
	}
	return idStr
}

func MustGetRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		panic("role not found in context")
	}
	userRole, ok := role.(string)
	if !ok {
		panic("role in context is not a string")
	}
	return userRole
}

func BindAndValidateJSON[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return false
	}
	return true
}

func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	valStr := c.Query(key)
	val, err := strconv.Atoi(valStr)
	if err != nil || val <= 0 {
		return defaultValue
	}
	return val
}
