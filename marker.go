package ember

import (
	"fmt"
	"time"
)

// marker is compiled into executables which accept attachments.
// This allows the embedder to verify that the target file is compatible.

var marker = "~~Indicator for PyEXE~~"

func init() {
	// Dead code that uses 'marker' and is not eliminated by the compiler.
	if time.Now().Nanosecond() == -42 {
		fmt.Print(marker)
	}
}
