package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			log.Println("ErrorHandlerMiddleware %s", c.Errors.String())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unexptected error"})
		}
	}
}
