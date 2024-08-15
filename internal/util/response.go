package util

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/model"
)

func SendSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.Set(constants.Response, model.Response{
		ResponseCode:    statusCode,
		ResponseMessage: message,
		ResponseData:    data,
	})
	ctx.Status(statusCode)
}

func SendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Set(constants.Response, model.Response{
		ResponseCode:    statusCode,
		ResponseMessage: constants.Error,
		ResponseError:   message,
	})
	ctx.Status(statusCode)
	ctx.Abort()
}

func Response(ctx *gin.Context, response model.ResponseInterface) {
	if response.Error != nil {
		SendErrorResponse(ctx, response.StatusCode, response.Message)
	} else {
		SendSuccessResponse(ctx, response.StatusCode, response.Message, response.Data)
	}
}
