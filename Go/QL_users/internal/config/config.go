package config
// viper: quản lý cấu hình, đọc từ file YAML, JSON
import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Service  string
	Username string
	Password string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

func Load() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	cfg.Server.Port = viper.GetString("server.port")

	cfg.Database.Host = viper.GetString("database.host")
	cfg.Database.Port = viper.GetInt("database.port")
	cfg.Database.Service = viper.GetString("database.service")
	cfg.Database.Username = viper.GetString("database.username")
	cfg.Database.Password = viper.GetString("database.password")

	cfg.JWT.Secret = viper.GetString("jwt.secret")
	cfg.JWT.ExpireHours = viper.GetInt("jwt.expireHours")

	return cfg, nil
}