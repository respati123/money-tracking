package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/respati123/money-tracking/docs"
	_ "github.com/respati123/money-tracking/docs"
	"github.com/respati123/money-tracking/internal/configs"
	"github.com/respati123/money-tracking/internal/configs/logger"
)

// @title My API
// @version 1.0
// @description Money Tracking Service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @produce application/json
// @consumes application/json

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func main() {

	config, viper := configs.InitConfig()
	log := logger.NewLogger(viper)
	db := configs.Database(config, log)
	redis := configs.NewRedis(viper, log)
	app := configs.NewGin(viper)

	// setup swagger
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	configs.Bootstrap(&configs.BootstrapConfig{
		DB:     db,
		Log:    log,
		Viper:  viper,
		App:    app,
		Config: config,
		Redis:  redis,
	})
	app.Run(":" + viper.GetString("PORT_SERVER"))
}
