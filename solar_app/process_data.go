package solar

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func process(c *gin.Context) {
	var data Request
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"detail": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
