package configs

import (
	"fmt"
	"time"

	"github.com/respati123/money-tracking/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database(config util.Config, log *logrus.Logger) *gorm.DB {

	postgresqlDbInfo1 := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(postgresqlDbInfo1), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(config.DB_POOL_IDLE)
	connection.SetMaxOpenConns(config.DB_POOL_MAX)
	connection.SetConnMaxLifetime(time.Second * time.Duration(config.DB_POOL_LIFETIME))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
