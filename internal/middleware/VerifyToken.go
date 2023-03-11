package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func VerifyToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		// Set example variable
		c.Set("example", "12345")
		fmt.Println("verifyToken")
		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}
