package configs

import (
	"fmt"

	"github.com/respati123/money-tracking/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(config util.Config) (db1 *gorm.DB, db2 *gorm.DB, err error) {

	fmt.Println(config.DB_NAME)

	postgresqlDbInfo1 := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
	)

	postgresqlDbInfo2 := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME_TEST,
	)

	connection1, err := gorm.Open(postgres.New(postgres.Config{
		DSN: postgresqlDbInfo1,
	}))

	connection2, err := gorm.Open(postgres.New(postgres.Config{
		DSN: postgresqlDbInfo2,
	}))

	return connection1, connection2, err
}
