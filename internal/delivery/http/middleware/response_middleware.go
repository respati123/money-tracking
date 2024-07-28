package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		var (
			response interface{}
		)

		value, exists := ctx.Get(constants.Response)

		if !exists {
			return
		}

		response = value

		switch ctx.Writer.Status() {
		case http.StatusOK:
			ctx.JSON(http.StatusOK, response)

		case http.StatusUnauthorized:
			ctx.JSON(http.StatusUnauthorized, response)

		case http.StatusBadRequest:
			ctx.JSON(http.StatusBadRequest, response)
		case http.StatusForbidden:
			ctx.JSON(http.StatusForbidden, response)
		case http.StatusCreated:
			ctx.JSON(http.StatusCreated, response)
		default:
			ctx.JSON(http.StatusInternalServerError, response)
		}

	}
}
