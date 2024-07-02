package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/respati123/money-tracking/docs"
	"github.com/respati123/money-tracking/internal/handlers"
	"github.com/respati123/money-tracking/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type RunServerParams struct {
	env *util.Config
}

func RunServer(env util.Config, db gorm.DB) {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

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

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + env.PORT_SERVER)

	versionOne := router.Group("/v1/")
	{
		authRouter := versionOne.Group("auth")
		{
			authRouter.POST("/login", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"success": true,
				})
			})
		}
	}
}
