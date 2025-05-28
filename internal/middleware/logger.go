package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logrus.WithFields(logrus.Fields{
			"status_code":  param.StatusCode,
			"latency":      param.Latency,
			"client_ip":    param.ClientIP,
			"method":       param.Method,
			"path":         param.Path,
			"error":        param.ErrorMessage,
			"body_size":    param.BodySize,
			"user_agent":   param.Request.UserAgent(),
			"timestamp":    param.TimeStamp.Format(time.RFC3339),
		}).Info("HTTP Request")
		
		return ""
	})
}

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logrus.WithFields(logrus.Fields{
			"error": recovered,
			"path":  c.Request.URL.Path,
			"method": c.Request.Method,
		}).Error("Panic recovered")
		
		c.JSON(500, gin.H{
			"code":    500,
			"message": "Internal server error",
		})
	})
}