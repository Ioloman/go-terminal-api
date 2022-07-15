package solar

import (
	"log"
	"net/http"

	"github.com/Ioloman/go-terminal-api/database"
	"github.com/gin-gonic/gin"
)

func process(c *gin.Context) {
	var data Request
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"detail": "error", "message": err.Error()})
		return
	}
	_, err := database.DB.NamedExec(
		`
		insert into
            processing.solar_panels_data
        set
            date = :date, person_id = :person_id,
			battery_voltage = :battery_voltage, solar_voltage = :solar_voltage,
			solar_current = :solar_current, load_current = :load_current
		`,
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}
