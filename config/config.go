package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Uri          string
	DatabaseName string
	Username     string
	Password     string
}

func (c *Config) Read() {
	viper.SetConfigType("yml")
	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Error decoding config, %v", err)
	}
}
