package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`

	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"DBname"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"database"`

	App struct {
		Version string `mapstructure:"version"`
	}

	Grpc struct {
		GRPCPort int `mapstructure:"grpcPort"`
	}
}

func LoadConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/app")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %w", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфигурации: %w", err)
	}

	return &cfg, nil
}
