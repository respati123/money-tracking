package configs

import (
	"fmt"

	"github.com/respati123/money-tracking/internal/util"
	"github.com/spf13/viper"
)

func InitConfig() (config util.Config, cfg *viper.Viper) {
	cfg = viper.New()
	cfg.SetConfigType("env")
	cfg.AddConfigPath(".")
	cfg.SetConfigName("app")
	err := cfg.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	errMarshal := cfg.Unmarshal(&config)

	if errMarshal != nil {
		fmt.Print(errMarshal.Error())
	}
	return config, cfg

}
