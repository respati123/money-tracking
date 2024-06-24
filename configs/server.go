package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/handlers"
	"github.com/respati123/money-tracking/util"
	"gorm.io/gorm"
)

type RunServerParams struct {
	env *util.Config
}

func RunServer(env util.Config, db gorm.DB) {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	version1 := router.Group("/v1")
	{
		authRouter := version1.Group("/auth")
		handlers.NewAuthHandler(authRouter, db, env)
	}

	router.Run(":" + env.PORT_SERVER)
	// create route here;

	// versionOne := router.Group("/v1/")
	// {
	// 	authRouter := versionOne.Group("auth")
	// 	{
	// 			authRouter.POST("/login", fun)
	// 	}
	// }
}
