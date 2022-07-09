package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	// read .env
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file is not located")
	}

	// setup logger
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	if os.Getenv("APP_ENV") == "prod" {
		path := filepath.Join(os.Getenv("LOG_DIR"), "general.log")
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err.Error())
		}
		log.SetOutput(file)
	}
}
