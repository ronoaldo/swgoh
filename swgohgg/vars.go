package swgohgg

import (
	"fmt"
	"regexp"
)

var (
	// ShapeNames is a list of all available mod shapes.
	ShapeNames = []string{"Transmitter", "Receiver", "Processor", "Holo-Array", "Data-Bus", "Multiplexer"}

	// StatNames is a list of all available mod stats.
	StatNames = []string{"Speed", "Critical Chance", "Critical Damage", "Critical Avoidance", "Accuracy", "Potency", "Tenacity", "Offense", "Defense", "Health", "Protection"}

	shapes = map[string]int{
		"Transmitter": 1,
		"Receiver":    2,
		"Processor":   3,
		"Holo-Array":  4,
		"Data-Bus":    5,
		"Multiplexer": 6,
	}

	errNotImplemented = fmt.Errorf("swgohgg: not implemented")

	nonDigits = regexp.MustCompile("[^0-9]")
)
