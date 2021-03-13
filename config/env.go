package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetENV() error {
	err := godotenv.Load("./env.local")
	if err != nil {
		fmt.Printf("can't read .envfile")
		return err
	}
	return nil
}
