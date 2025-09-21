package config

import "os"

type Config struct {
	Addr         string
	DbConfig     DbConfig
	JWTSecretKey string
	Version      string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Addr:         getEnv("ADDR", ":8080"),
		Version:      getEnv("VERSION", "0.0.1"),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "secret"),
		DbConfig: DbConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			DbName:   getEnv("DB_NAME", "basic_http"),
			User:     getEnv("DB_USER", "basic_http"),
			Password: getEnv("DB_PASSWORD", "password"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

