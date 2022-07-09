package solar

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func process(c *gin.Context) {
	defer c.Request.Body.Close()
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"detail": "error", "message": "couldn't read request body"})
		return
	}

	var data gin.H
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("json decoding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"detail": "error", "message": "json decoding error"})
		return
	}

	_, ok1 := data["terminal_id"]
	_, ok2 := data["date"]
	if !ok1 || !ok2 {
		log.Println("validation error")
		c.JSON(http.StatusBadRequest, gin.H{"detail": "error", "message": "required field 'date' or 'terminal_id' is missing"})
		return
	}

	c.JSON(http.StatusOK, data)
}
