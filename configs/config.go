package configs

import (
	"fmt"

	"github.com/respati123/money-tracking/util"
	"github.com/spf13/viper"
)

func InitConfig(path string) (config util.Config, errConfig error) {
	cfg := viper.New()
	cfg.SetConfigType("env")
	cfg.AddConfigPath(path)
	cfg.SetConfigName("app")

	err := cfg.ReadInConfig()

	if err != nil {
		fmt.Printf("error config: %s", err.Error())
	}

	errMarshal := cfg.Unmarshal(&config)

	if errMarshal != nil {
		fmt.Print(errMarshal.Error())
	}
	return config, nil

}
