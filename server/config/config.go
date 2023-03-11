package config

import "os"

type appConfig struct {
	PostgresInfo *PostgresInfo
	ServerInfo   *ServerInfo
}

type PostgresInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Timezone string
	Sslmode  string
}

type ServerInfo struct {
	Host     string
	Port     string
	Domain   string
	Protocol string
}

var AppConfig *appConfig

func init() {
	AppConfig = LoadConfig()
}

func LoadConfig() *appConfig {
	appConfig := &appConfig{
		PostgresInfo: &PostgresInfo{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Dbname:   os.Getenv("POSTGRES_DB"),
			Timezone: os.Getenv("POSTGRES_TIME_ZONE"),
			Sslmode:  os.Getenv("POSTGRES_SSL_MODE"),
		},
		ServerInfo: &ServerInfo{
			Host:     "localhost",
			Port:     "8080",
			Domain:   "localhost:8080",
			Protocol: "http",
		},
	}
	return appConfig
}

func GetPostgresInfo() *PostgresInfo {
	return AppConfig.PostgresInfo
}

func GetServerInfo() *ServerInfo {
	return AppConfig.ServerInfo
}
