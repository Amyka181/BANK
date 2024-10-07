package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadEnv() (*Config, error) {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Ошибка загрузки .env файла: %v", err)
	}

	// Читаем переменные окружения
	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return config, nil
}

func ConnectDB(cfg *Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
