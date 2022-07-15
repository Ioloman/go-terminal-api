package solar

import (
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

type Request struct {
	TerminalId     int          `json:"terminal_id" db:"person_id" binding:"required"`
	Date           time.Time    `json:"date" db:"date" binding:"required"`
	BatteryVoltage encodedValue `json:"battery_voltage" db:"battery_voltage"`
	SolarVoltage   encodedValue `json:"solar_voltage" db:"solar_voltage"`
	SolarCurrent   encodedValue `json:"solar_current" db:"solar_current"`
	LoadCurrent    encodedValue `json:"load_current" db:"load_current"`
}

// parse number that was encoded in signed hex with length of 4
func parseSigned(i int32) int32 {
	mask := int32(1 << 15)
	return -(i & mask) | (i & (mask - 1))
}

// type for value that comes encoded in a special way.
// incoming data is a string of hex-encoded bytes and result comes in position of 4th and 5th bytes.
// in "010402053D7BB1" 053D is desired signed (15th bit) number == 1341
type encodedValue float64

func (v *encodedValue) UnmarshalJSON(b []byte) error {
	// extract hex string
	var strValue string
	if err := json.Unmarshal(b, &strValue); err != nil {
		return err
	}

	i, err := strconv.ParseInt(strValue[6:10], 16, 32)
	if err != nil {
		return err
	}

	*v = encodedValue(float64(parseSigned(int32(i))) / 100)

	return nil
}
