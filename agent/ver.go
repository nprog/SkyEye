package main

import (
	"fmt"
)

const (
	VERSION = "0.01"
)

func version() {
	fmt.Printf("\n--| SkyEye Agent | version: %s |--\n\n", VERSION)
}
