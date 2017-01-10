package swgohgg

import (
	"fmt"
)

var (
	// ShapeNames
	ShapeNames = []string{"Transmitter", "Receiver", "Processor", "Holo-Array", "Data-Bus", "Multiplexer"}

	shapes = map[string]int{
		"Transmitter": 1,
		"Receiver":    2,
		"Processor":   3,
		"Holo-Array":  4,
		"Data-Bus":    5,
		"Multiplexer": 6,
	}

	errNotImplemented = fmt.Errorf("swgohgg: not implemented")
)
