package middleware

import (
    "sync/atomic"

    "github.com/gin-gonic/gin"
)

var reqCount uint64

// RequestCounter returns a middleware that increments an internal request counter.
func RequestCounter() gin.HandlerFunc {
    return func(c *gin.Context) {
        atomic.AddUint64(&reqCount, 1)
        c.Next()
    }
}

// GetRequestCount returns the current total request count.
func GetRequestCount() uint64 {
    return atomic.LoadUint64(&reqCount)
}
