package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/delivery/http"
	"github.com/respati123/money-tracking/internal/delivery/http/middleware"
	"github.com/respati123/money-tracking/internal/delivery/http/route"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Viper  *viper.Viper
	App    *gin.Engine
	Config util.Config
}

func Bootstrap(config *BootstrapConfig) {

	// setup repository
	userRepository := repository.NewUserRepository(config.DB, config.Log)
	authRepository := repository.NewAuthRepository(config.DB, config.Log)

	// setup service
	userUseCase := usecase.NewUserService(userRepository, config.DB, config.Log)
	authUseCase := usecase.NewAuthUsecase(config.Log, authRepository, userRepository, config.Config)

	// setup controllers
	userController := http.NewUserController(userUseCase, config.DB, config.Log)
	authController := http.NewAuthController(config.Log, authUseCase)

	// setup middleware
	traceIdMiddleware := middleware.NewTraceMiddleware()

	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		AuthController:    authController,
		TraceIdMiddleware: traceIdMiddleware,
	}
	routeConfig.Setup()
}
