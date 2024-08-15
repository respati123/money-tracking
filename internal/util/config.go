package util

type Config struct {
	Env string `mapstructure:"env"`

	// DB
	DB_HOST          string `mapstructure:"DB_HOST"`
	DB_USER          string `mapstructure:"DB_USER"`
	DB_PASS          string `mapstructure:"DB_PASS"`
	DB_NAME          string `mapstructure:"DB_NAME"`
	DB_NAME_TEST     string `mapstructure:"DB_NAME_TEST"`
	DB_PORT          string `mapstructure:"DB_PORT"`
	DB_POOL_IDLE     int    `mapstructure:"DB_POOL_IDLE"`
	DB_POOL_MAX      int    `mapstructure:"DB_POOL_MAX"`
	DB_POOL_LIFETIME int    `mapstructure:"DB_POOL_LIFETIME"`

	// JWT
	JWT_SECRET_KEY          string `mapstructure:"JWT_SECRET_KEY"`
	JWT_EXPIRE_TIME         int    `mapstructure:"JWT_EXPIRE_TIME"`
	JWT_EXPIRE_REFRESH_TIME int    `mapstructure:"JWT_EXPIRE_REFRESH_TIME"`

	// PORT SERVER
	PORT_SERVER string `mapstructure:"PORT_SERVER"`
}
