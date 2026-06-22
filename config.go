package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Cors struct {
		AllowedOrigins []string `mapstructure:"allowed_origins"`
	} `mapstructure:"cors"`
}

var Cfg Config

func Load(path string) error {
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("cors.allowed_origins", []string{"*"})

	viper.SetEnvPrefix("APP")
	viper.BindEnv("server.port", "APP_SERVER_PORT")
	viper.BindEnv("cors.allowed_origins", "APP_CORS_ALLOWED_ORIGINS")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	p := viper.Get("APP_SERVER_PORT")
	println(fmt.Sprintf("port:%v", p))

	if port := viper.GetInt("APP_SERVER_PORT"); port != 0 {
		Cfg.Server.Port = port
	}

	if origins := viper.GetStringSlice("APP_CORS_ALLOWED_ORIGINS"); len(origins) > 0 {
		Cfg.Cors.AllowedOrigins = origins
	}

	return nil
}
