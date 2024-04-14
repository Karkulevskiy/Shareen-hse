package configs

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(filename string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(filename)
	config.AddConfigPath(".\\internal\\configs\\")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error while parsing configuraion file\n", err)
	}
	return config
}
