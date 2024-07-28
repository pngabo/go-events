package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		log.Printf("Request: %s %s took %v", ctx.Request.Method, ctx.Request.URL.Path, duration)
	}
}
