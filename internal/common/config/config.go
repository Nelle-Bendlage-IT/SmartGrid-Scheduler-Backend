package config

import (
	"github.com/JeremyLoy/config"
)

type Config struct {
	Supabase struct {
		URL string `config:"URL"`
		Key string `config:"KEY"`
	}
	Port            int
	CorrentlyAPIKey string `config:"CORRENTLY_API_KEY"`
	DB              struct {
		User string `config:"USER"`
		Pass string `config:"PASS"`
		URL  string `config:"URL"`
	}
}

var appConfig Config

func init() {
	err := config.FromEnv().To(&appConfig)
	if err != nil {
		panic(err)
	}

	err = config.FromEnv().Sub(&appConfig.Supabase, "SUPABASE")
	if err != nil {
		panic(err)
	}

	if appConfig.Port == 0 {
		appConfig.Port = 4700
	}
}

func GetConfig() Config {
	return appConfig
}
