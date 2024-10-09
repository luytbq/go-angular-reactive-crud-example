package config

import (
	"os"

	"github.com/joho/godotenv"
)

var App Config

type Config struct {
	SERVER_PORT       string
	SERVER_API_PREFIX string

	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_PASSWORD string
	PG_DB_NAME  string
	PG_SSL_MODE string
}

const (
	SERVER_PORT       = "SERVER_PORT"
	SERVER_API_PREFIX = "SERVER_API_PREFIX"

	PG_HOST     = "PG_HOST"
	PG_PORT     = "PG_PORT"
	PG_USER     = "PG_USER"
	PG_PASSWORD = "PG_PASSWORD"
	PG_DB_NAME  = "PG_DB_NAME"
	PG_SSL_MODE = "PG_SSL_MODE"
)

func init() {
	_ = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	App.SERVER_PORT = getEnv(SERVER_PORT, ":8080")
	App.SERVER_API_PREFIX = getEnv(SERVER_API_PREFIX, "/api/v1")
	App.PG_HOST = getEnv(PG_HOST, "127.0.0.1")
	App.PG_PORT = getEnv(PG_PORT, "5432")
	App.PG_USER = getEnv(PG_USER, "postgres")
	App.PG_PASSWORD = getEnv(PG_PASSWORD, "postgres")
	App.PG_DB_NAME = getEnv(PG_PASSWORD, "postgres")
	App.PG_SSL_MODE = getEnv(PG_SSL_MODE, "disable")
}

func getEnv(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return def
	}
}
