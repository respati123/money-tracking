package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewGin(config *viper.Viper) *gin.Engine {
	var app = gin.New()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "connected",
		})
	})
	return app

}
