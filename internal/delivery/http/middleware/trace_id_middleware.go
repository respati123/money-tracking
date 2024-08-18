package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/respati123/money-tracking/internal/constants"
)

const TraceIDHeader = "x-trace-id"

func NewTraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader(TraceIDHeader)

		if traceID == "" {
			traceID = uuid.New().String()
		}

		ctx.Set(constants.TraceID, traceID)
		ctx.Request = ctx.Request.WithContext(ctx)
		ctx.Writer.Header().Set(TraceIDHeader, traceID)
		ctx.Next()

	}
}
