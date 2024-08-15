package configs

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/delivery/http/middleware"
	"github.com/spf13/viper"
)

func NewGin(config *viper.Viper) *gin.Engine {
	var app = gin.New()
	app.Use(middleware.TimeoutMiddleware(time.Duration(config.GetInt("TIMEOUT"))))
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "connected",
		})
	})
	return app

}
