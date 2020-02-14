package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	return cors.New(config)
}

