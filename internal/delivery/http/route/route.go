package route

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/delivery/http"
	"github.com/respati123/money-tracking/internal/delivery/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController http.UserController
	AuthController http.AuthController
	//
	TraceIdMiddleware gin.HandlerFunc
}

func (c *RouteConfig) Setup() {
	c.App.Use(middleware.NewTraceMiddleware())
	c.App.Use(middleware.ResponseMiddleware())
	c.SetupPublicRoute()
	c.SetupPrivateRoute()
	c.SetupSwagger()
}

func (c *RouteConfig) SetupSwagger() {
	c.App.GET("api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (c *RouteConfig) SetupPublicRoute() {
	c.App.POST("/api/v1/auth/login", c.AuthController.Login)
	c.App.POST("/api/v1/auth/register", c.AuthController.Register)
}

func (c *RouteConfig) SetupPrivateRoute() {
	c.App.GET("/api/v1/users/:user_code", c.UserController.GetUser)
	c.App.POST("/api/v1/users", c.UserController.CreateUser)
	c.App.PUT("/api/v1/users/:uuid", c.UserController.UpdateUser)
	c.App.DELETE("/api/v1/users/:uuid", c.UserController.DeleteUser)
}
