package main

import (
	"github.com/woshikedayaa/cl/cmd/cl"

	// commands
	_ "github.com/woshikedayaa/cl/cmd/cl/in"
	_ "github.com/woshikedayaa/cl/cmd/cl/version"
)

func main() {
	cl.Main()
}
