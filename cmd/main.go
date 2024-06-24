package main

import (
	"fmt"
	"log"

	"github.com/respati123/money-tracking/configs"
	"github.com/respati123/money-tracking/internal/database"
)

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
