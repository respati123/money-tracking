package main

import (
	"github.com/respati123/money-tracking/configs"
	"github.com/respati123/money-tracking/configs/logger"
	"github.com/respati123/money-tracking/docs"
	_ "github.com/respati123/money-tracking/docs"
)

// @title My API
// @version 1.0
// @description Money Tracking Service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api/v1
func main() {

	config, viper := configs.InitConfig()
	log := logger.NewLogger(viper)
	db := configs.Database(config, log)
	redis := configs.NewRedis(viper, log)
	app := configs.NewGin(viper)

	// setup swagger
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
