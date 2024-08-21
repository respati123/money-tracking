package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/respati123/money-tracking/internal/delivery/http"
	"github.com/respati123/money-tracking/internal/delivery/http/middleware"
	"github.com/respati123/money-tracking/internal/delivery/http/route"
	"github.com/respati123/money-tracking/internal/model/converter"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Log    *zap.Logger
	Viper  *viper.Viper
	App    *gin.Engine
	Config util.Config
	Redis  *redis.Client
}

func Bootstrap(config *BootstrapConfig) {

	//
	converter := &converter.Converter{
		UserConverter: converter.NewUserConverter(),
	}

	// setup repository
	userRepository := repository.NewUserRepository(config.Log)
	authRepository := repository.NewAuthRepository(config.Log)
	roleRepository := repository.NewRoleRepository(config.Log)
	categoryRepository := repository.NewCategoryRepository(config.Log)
	transactionRepository := repository.NewTransactionRepository(config.Log)

	// setup service
	userUseCase := usecase.NewUserUsecase(config.DB, config.Log, converter, userRepository)
	authUseCase := usecase.NewAuthUsecase(config.DB, config.Log, config.Config, config.Redis, authRepository, userRepository)
	roleUseCase := usecase.NewRoleUsecase(config.DB, config.Log, converter, roleRepository)
	categoryUsecase := usecase.NewCategoryUsecase(config.DB, config.Log, converter, categoryRepository)
	transactionUsecase := usecase.NewTransactionUsecase(config.DB, config.Log, converter, transactionRepository)

	// setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	authController := http.NewAuthController(authUseCase, config.Log)
	roleController := http.NewRoleController(roleUseCase, config.Log)
	categoryController := http.NewCategoryController(categoryUsecase, config.Log)
	transactionController := http.NewTransactionController(transactionUsecase, config.Log)

	// setup middleware
	traceIdMiddleware := middleware.NewTraceMiddleware()
	responseMiddleware := middleware.ResponseMiddleware()
	authMiddleware := middleware.AuthMiddleware(config.Redis, config.Viper, config.Log, config.DB)

	routeConfig := route.RouteConfig{
		App: config.App,

		// controller
		UserController:        userController,
		AuthController:        authController,
		RoleController:        roleController,
		CategoryController:    categoryController,
		TransactionController: transactionController,

		// middleware
		TraceIdMiddleware:  traceIdMiddleware,
		ResponseMiddleware: responseMiddleware,
		AuthMiddleware:     authMiddleware,
	}
	routeConfig.Setup()
}
