package main

import (
	"autossh/src/app"
)

var (
	Version = "1.0"
	Build   = "1.0"
)

func main() {
	app.Version = Version
	app.Build = Build
	app.Run()
}
