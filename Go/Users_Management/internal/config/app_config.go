package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database-oracle"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	SMTP    SMTPConfig    `mapstructure:"smtp"`
    Frontend FrontendConfig `mapstructure:"frontend"`
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

type SMTPConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    From     string `mapstructure:"from"`
}

type FrontendConfig struct {
    URL string `mapstructure:"url"`
}