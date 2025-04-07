package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			fmt.Println("ErrorHandlerMiddleware" + c.Errors.String())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unexptected error"})
		}
	}
}
