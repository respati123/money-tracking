package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		context, cancel := context.WithTimeout(c.Request.Context(), timeout*time.Second)

		defer cancel()

		c.Request = c.Request.WithContext(context)

		finished := make(chan struct{})
		panicChan := make(chan interface{})

		go func() {
			defer close(finished)
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			c.Next()
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-finished:
			return
		case <-context.Done():
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "request timeout"})
			return
		}
	}
}
