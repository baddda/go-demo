package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("\n[%s] %s %s %v\n", c.Request.Method, c.Request.URL.Path, c.ClientIP(), latency)
	}
}
