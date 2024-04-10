// package config contains the methods with extract values of enverioment variables
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	STRING_MYSQL_CONNECTION = ""
	API_PORT                = ""
	JWT_SECRET_KEY          []byte
)

func LoadEnvs() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	return nil
}

func GetEnv() {

	LoadEnvs()

	API_PORT = ":" + os.Getenv("API_PORT")

	fmt.Println(os.Getenv("API_PORT"))

	STRING_MYSQL_CONNECTION = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER_NAME"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRETE_KEY"))
}
