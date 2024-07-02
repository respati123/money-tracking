package main

import (
	"fmt"
	"log"

	"github.com/respati123/money-tracking/configs"
	_ "github.com/respati123/money-tracking/docs"
	"github.com/respati123/money-tracking/internal/database"
)

// @title My API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {

	config, err := configs.InitConfig(".")

	if err != nil {
		defer log.Fatal("env not found", err.Error())
	}

	dbReal, dbTest, err := configs.Database(config)
	database.NewMigration(dbReal, dbTest)

	if err != nil {
		defer fmt.Printf("error database %s", err.Error())
	}

	configs.RunServer(config, *dbReal)
}
