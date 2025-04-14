package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func ServerTiming() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		dbStart := time.Now()
		c.Set("dbStart", dbStart)

		c.Next()

		dbDuration := time.Since(dbStart).Milliseconds()
		totalDuration := time.Since(start).Milliseconds()

		c.Header("Server-Timing",
			fmt.Sprintf("db;dur=%d, total;dur=%d", dbDuration, totalDuration))
		c.Header("Timing-Allow-Origin", "*")
	}
}
