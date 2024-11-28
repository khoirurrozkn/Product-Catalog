package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Driver   string
}

type Config struct {
	DbConfig
}

func (c *Config) ReadConfig() error {

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error load .env file")
	}

	c.DbConfig = DbConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Driver: os.Getenv("DB_DRIVER"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.DbName == "" || 
		c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" {

		return fmt.Errorf("missing env")
	}

	return nil
}

func NewConfig() ( *Config, error ) {
	cfg := &Config{}

	if err := cfg.ReadConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}