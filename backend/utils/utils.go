package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserIDFromContext extracts the user ID from the Gin context (assuming it is stored as a string in context)
func GetUserIDFromContext(c *gin.Context) int {
	userIDStr, exists := c.Get("userID")
	if !exists {
		return 0
	}
	switch v := userIDStr.(type) {
	case string:
		id, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return id
	case int:
		return v
	default:
		return 0
	}
}
