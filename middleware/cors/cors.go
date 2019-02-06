package cors

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"time"
	"strconv"
)

// CORS func
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		MaxAge := 6 * time.Hour

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Request-Methods", "POST, GET")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-PINGOTHER")
		c.Header("Access-Control-Max-Age", strconv.Itoa(int(MaxAge)))
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}