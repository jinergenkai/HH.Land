package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware xử lý CORS đúng cách
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin") // Lấy Origin từ request
		if origin != "" {
			// Chỉ thêm headers nếu có Origin (tránh thêm headers không cần thiết)
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
			c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Vary", "Origin")

			// Nếu là Preflight request (OPTIONS), trả về ngay mà không xử lý tiếp
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}

		// Tiếp tục request bình thường
		c.Next()
	}
}
