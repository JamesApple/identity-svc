package cfg

import (
	"os"
)

var configs = Configs{
	"test": &Config{
		postgresURL: "postgres://localhost/identity_test?sslmode=disable",
	},
	"development": &Config{
		postgresURL: "postgres://localhost/identity?sslmode=disable",
	},
}

type Configs map[string]*Config

type Config struct {
	postgresURL string
}

func MakeConfig() *Config {
	return configs[getEnv()]
}

func getEnv() string {
	env, ok := os.LookupEnv("GO_ENV")
	if ok {
		return env
	}
	return "development"
}
