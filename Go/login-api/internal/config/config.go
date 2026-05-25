package config

import "os"

type Config struct {
	Port string
	JWTSecret string
}

func Load() Config {
	return Config{
		Port: os.Getenv("APP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}