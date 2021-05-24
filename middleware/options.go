package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}
