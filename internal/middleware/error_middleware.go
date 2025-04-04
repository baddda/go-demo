package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("ErrorHandlerMiddleware" + c.Errors.String())
		if len(c.Errors) > 0 {
			// Don't show the error to outside world.
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unexptected error"})
		}
	}
}
