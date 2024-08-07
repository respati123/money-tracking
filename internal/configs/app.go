package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/delivery/http"
	"github.com/respati123/money-tracking/internal/delivery/http/middleware"
	"github.com/respati123/money-tracking/internal/delivery/http/route"
	"github.com/respati123/money-tracking/internal/model/converter"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Log    *logger.CustomLogger
	Viper  *viper.Viper
	App    *gin.Engine
	Config util.Config
}

func Bootstrap(config *BootstrapConfig) {

	//
	converter := &converter.Converter{
		UserConverter: converter.NewUserConverter(),
	}

	// setup repository
	userRepository := repository.NewUserRepository(config.Log)
	authRepository := repository.NewAuthRepository(config.Log)

	// setup service
	userUseCase := usecase.NewUserUsecase(config.DB, config.Log, converter, userRepository)
	authUseCase := usecase.NewAuthUsecase(config.DB, config.Log, config.Config, authRepository, userRepository)

	// setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	authController := http.NewAuthController(authUseCase, config.Log)

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
