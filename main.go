package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Getenv("LOG_DIR"))
}
