package config

type Config struct {
	Server     ServerConfig
	PostgreSql PostgreSqlConfig
	Security   SecurityConfig
}

type ServerConfig struct {
	Port    int
	Version string
}

type PostgreSqlConfig struct {
	Server   string
	Port     int
	Database string
	User     string
	Password string
}

type SecurityConfig struct {
	JWTSecret string
}
