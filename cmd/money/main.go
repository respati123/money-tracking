package main

import (
	"github.com/respati123/money-tracking/configs"
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
	log := configs.NewLogger(viper)
	db := configs.Database(config, log)
	app := configs.NewGin(viper)

	// setup swagger
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	configs.Bootstrap(&configs.BootstrapConfig{
		DB:     db,
		Log:    log,
		Viper:  viper,
		App:    app,
		Config: config,
	})
	app.Run(":" + viper.GetString("PORT_SERVER"))
}
