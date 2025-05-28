package middleware

import (
	"time"

	"github.com/dengxing/gin-grom/internal/config"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	cfg := config.GetConfig()
	
	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// 检查是否允许该来源
		allowed := false
		for _, allowedOrigin := range cfg.CORS.AllowOrigins {
			if allowedOrigin == "*" || allowedOrigin == origin {
				allowed = true
				break
			}
		}
		
		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		
		c.Header("Access-Control-Allow-Methods", joinStrings(cfg.CORS.AllowMethods, ", "))
		c.Header("Access-Control-Allow-Headers", joinStrings(cfg.CORS.AllowHeaders, ", "))
		c.Header("Access-Control-Expose-Headers", joinStrings(cfg.CORS.ExposeHeaders, ", "))
		
		if cfg.CORS.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		
		if cfg.CORS.MaxAge > 0 {
			c.Header("Access-Control-Max-Age", string(rune(cfg.CORS.MaxAge*int(time.Hour.Seconds()))))
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}