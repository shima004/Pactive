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
			Host:     os.Getenv("SERVER_HOST"),
			Port:     "8080",
			Domain:   "shimapaca.net",
			Protocol: os.Getenv("SERVER_PROTOCOL"),
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
