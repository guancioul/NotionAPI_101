package util

import (
	"github.com/spf13/viper"
)

type ConfigHandler struct {
}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{}
}

func (c *ConfigHandler) GetSecretConfig() *viper.Viper {
	viper.AddConfigPath("./config") // config所在的目錄路徑
	viper.SetConfigName("secret")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetViper()
}
