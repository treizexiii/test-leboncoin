package main

import "fmt"
import "github.com/spf13/viper"

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
	viper.SetConfigName("config") // nom sans extension
	// on va faire un config.yaml ou yml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path) // ex: "."
	viper.AddConfigPath(".")  // fallback

	// Defaults
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("cors.allowed_origins", []string{"*"})

	// Env: APP_SERVER_PORT, APP_CORS_ALLOWED_ORIGINS (liste séparée par virgules si tu veux aller plus loin)
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv() // les env écrasent le fichier[web:41][web:49]

	if err := viper.ReadInConfig(); err != nil {
		// Pas grave si le fichier manque, on garde defaults + env
		// tu peux choisir de le rendre fatal si tu préfères
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
