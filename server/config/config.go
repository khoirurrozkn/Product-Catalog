package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

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
	TokenConfig
}

type TokenConfig struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
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

	token_lifetime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return errors.New("fail parse token life time")
	}

	pemKey, err := os.ReadFile("ec-secp256k1-priv-key.pem")
	if err != nil {
		return errors.New("cannot read .pem key")
	}

	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("ISSUER_NAME"),
		JwtSignatureKey: pemKey,
		JwtLifeTime:     time.Duration(token_lifetime) * time.Hour,
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