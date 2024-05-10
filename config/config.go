package config

import (
	"fmt"
	"github.com/Lincyaw/PaperGraph-backend/logger"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile("config.toml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	logger.Debug("initialized config", "config", viper.AllSettings())
	logger.Global.Debug("initialized config", "config", viper.AllSettings())
}
