package config

import "os"

type Config struct {
	Port string
	JWTSecret string
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	DBService string
}
// Đọc các biến môi trường và trả về struct Config
func Load() Config {

	return Config{
		Port: os.Getenv("APP_PORT"),

		JWTSecret: os.Getenv("JWT_SECRET"),

		DBUser: os.Getenv("DB_USER"),

		DBPassword: os.Getenv("DB_PASSWORD"),

		DBHost: os.Getenv("DB_HOST"),

		DBPort: os.Getenv("DB_PORT"),

		DBService: os.Getenv("DB_SERVICE"),
	}
}