package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return "", errors.New("dburl is not set")
	}

	return dbURL, nil

}
