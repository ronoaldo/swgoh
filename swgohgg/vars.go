package swgohgg

import (
	"fmt"
)

var (
	ShapeNames = []string{"Transmitter", "Receiver", "Processor", "Holo-Array", "Data-Bus", "Multiplexer"}

	errNotImplemented = fmt.Errorf("swgohgg: not implemented")
)
