package configs

import (
	"fmt"
	"time"

	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Database(config util.Config, log *zap.Logger) *gorm.DB {

	postgresqlDbInfo1 := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(postgresqlDbInfo1), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.New(zap.NewStdLog(log), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatal("failed to connect database: %v", zap.Error(err))
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect database: %v", zap.Error(err))
	}

	connection.SetMaxIdleConns(config.DB_POOL_IDLE)
	connection.SetMaxOpenConns(config.DB_POOL_MAX)
	connection.SetConnMaxLifetime(time.Second * time.Duration(config.DB_POOL_LIFETIME))

	return db
}
