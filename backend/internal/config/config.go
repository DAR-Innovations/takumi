package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env          string `mapstructure:"ENV"`
	Port         string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	ServerUrl    string `mapstructure:"SERVER_URL"`
	FrontendUrl  string `mapstructure:"FRONTEND_URL"`
}

var cfg *Config

func LoadConfig(file string) (*Config, error) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	cfg = &config
	return cfg, nil
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("config not loaded")
	}
	return cfg
}
