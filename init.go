package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getLogFile(name string) *os.File {
	path := filepath.Join(os.Getenv("LOG_DIR"), name)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err.Error())
	}
	return file
}

func init() {
	// read .env
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file is not located")
	}

	// config gin
	gin.SetMode(os.Getenv("GIN_MODE"))
	io.MultiWriter()
	if os.Getenv("APP_ENV") == "prod" {
		file := getLogFile("gin.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}

	// setup logger
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	if os.Getenv("APP_ENV") == "prod" {
		file := getLogFile("general.log")
		log.SetOutput(file)
	}
}
