package main

import (
	solar "github.com/Ioloman/go-terminal-api/solar_app"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/Ioloman/go-terminal-api/utils/middleware"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression, gzip.WithDecompressFn(middleware.DecompressFn)))

	solar.RunApp(r)

	r.Run(":8000")
}
