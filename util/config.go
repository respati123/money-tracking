package util

type Config struct {
	Env          string `mapstructure:"env"`
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_PASS      string `mapstructure:"DB_PASS"`
	DB_NAME      string `mapstructure:"DB_NAME"`
	DB_NAME_TEST string `mapstructure:"DB_NAME_TEST"`
	DB_PORT      string `mapstructure:"DB_PORT"`
	PORT_SERVER  string `mapstructure:"PORT_SERVER"`
}
