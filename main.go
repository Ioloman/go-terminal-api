package main

import (
	solar "github.com/Ioloman/go-terminal-api/solar_app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	solar.RunApp(r)

	r.Run(":8000")
}
