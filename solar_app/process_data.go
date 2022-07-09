package solar

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func process(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
