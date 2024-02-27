package middlewares

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(interval time.Duration, limit int) gin.HandlerFunc {
    bucket := ratelimit.NewBucketWithQuantum(interval, int64(limit), int64(limit))
    return func(c *gin.Context) {
        if bucket.TakeAvailable(1) == 0 {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
            c.Abort()
            return
        }
        c.Next()
    }
}

