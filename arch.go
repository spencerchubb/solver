package solver

import (
	"fmt"
	"strconv"
)

func check32Bit() {
	if strconv.IntSize == 32 {
		fmt.Println("WARNING: You are using a version of Go with 32-bit architecture. We recommend using a version of Go with 64-bit architecture, because it may run out of memory with 32-bit.")
	}
}
