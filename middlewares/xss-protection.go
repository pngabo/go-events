package middlewares

import (
	"github.com/gin-gonic/gin"
)

// XSSProtectionMiddleware adds XSS protection headers
func XSSProtectionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		ctx.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		ctx.Next()
	}
}
