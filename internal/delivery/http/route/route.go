package route

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/delivery/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController *http.UserController
	AuthController *http.AuthController
	RoleController *http.RoleController

	// middleware
	TraceIdMiddleware  gin.HandlerFunc
	ResponseMiddleware gin.HandlerFunc
	AuthMiddleware     gin.HandlerFunc
}

func (c *RouteConfig) Setup() {
	c.App.Use(c.TraceIdMiddleware)
	c.App.Use(c.ResponseMiddleware)
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
	protected := c.App.Group("")
	protected.Use(c.AuthMiddleware)

	users := protected.Group("/api/v1/users")
	users.POST("/list", c.UserController.GetListUser)
	users.POST("/", c.UserController.CreateUser)
	users.DELETE("/:uuid", c.UserController.Delete)
	users.PUT("/:uuid", c.UserController.Update)
	users.GET("/:user_code", c.UserController.GetUser)

	roles := protected.Group("/api/v1/roles")
	{
		roles.POST("/", c.RoleController.Create)
	}
}
