package config

import "github.com/spf13/viper"

type Config struct {
	DbUser string `mapstructure:"DBUSER"`
	DbPass string `mapstructure:"DBPASS"`
	DbIp   string `mapstructure:"DBIP"`
	DbName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
}

var LocalConfig *Config

func initConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
func SetConfig() {
	LocalConfig = initConfig()
}
