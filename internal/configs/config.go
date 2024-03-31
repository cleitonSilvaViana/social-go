// package config contains the methods with extract values of enverioment variables
package configs

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type envs struct {
	SGBD_MYSQL                   string `env:"SGBD_MYSQL,required"`
	MYSQL_DATABASE          string `env:"MYSQL_DATABASE"`
	MYSQL_USER_NAME     string `env:"MYSQL_USER_NAME"`
	MYSQL_USER_PASSWORD string `env:"MYSQL_USER_PASSWORD"`

	SGBD_MONGODB                   string `env:"SGBD_MONGO,required"`
	MONGO_DATABASE          string `env:"MONGO_DATABASE"`
	MONGO_USER_NAME     string `env:"MONGO_USER_NAME"`
	MONGO_USER_PASSWORD string `env:"MONGO_USER_PASSWORD"`

	API_PORT string `env:"API_PORT,required"`
}

var ENV envs

const env_file = ".env"

// LoadEnvFile realize the parse in value of enverioment variables to an struct on type envs.
// This struct contaisn the configs necessary to init application.
func loadEnvFiles() error {
	_, err := os.Stat((env_file))
	if err != nil {
		return err
	}
	return godotenv.Load(env_file)
}

// GetConfing get values of enverionment variables and set a global variable.
// The values in enverionment variables are necessary to start application.
func GetConfig() error {
	err := loadEnvFiles()
	if err != nil {
		return err
	}

	err = env.Parse(&ENV)
	if err != nil {
		return err
	}

	return nil
}
