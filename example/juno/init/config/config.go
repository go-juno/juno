package config

import (
	"log"

	"juno/init/flag"

	"github.com/spf13/viper"
)

var Config *Configuration

// Server
type Server struct {
	Grpc Grpc `yaml:"grpc"`
	Http Http `yaml:"http"`
}

// Grpc
type Grpc struct {
	Port int `yaml:"port"`
}

// Http
type Http struct {
	Port int `yaml:"port"`
}

// Database
type Database struct {
	Mysql Mysql `yaml:"mysql"`
}

// Mysql
type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
}

// Yaml2Go
type Configuration struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func init() {
	viper.SetConfigFile(flag.Flagconf)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("config file: %s : %v", flag.Flagconf, err)
	}
	log.Printf("config file is %s", flag.Flagconf)
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unmarshal err: %s : %v", flag.Flagconf, err)
	}
}
