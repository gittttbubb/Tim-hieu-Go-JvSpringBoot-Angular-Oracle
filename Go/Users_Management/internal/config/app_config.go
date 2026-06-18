package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"databse-oracle"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	ServiceName string `mapstructure:"service_name"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret              string `mapstructure:"secret"`
	AccessExpiryMinutes int    `mapstructure:"access_expiry_second"`
	RefreshExpiryHours  int    `mapstructure:"refresh_expiry_hours"`
}