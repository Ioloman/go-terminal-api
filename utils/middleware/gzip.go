package middleware

import (
	"compress/gzip"

	"github.com/gin-gonic/gin"
)

func DecompressFn(c *gin.Context) {
	reader, _ := gzip.NewReader(c.Request.Body)
	c.Request.Body = reader
}
