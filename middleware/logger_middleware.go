package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	start := time.Now()

	c.Next()

	elapsed := time.Since(start)

	//logs
	fmt.Printf("%s %s %s %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP(), elapsed)
}
