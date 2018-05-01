package main

import (
	"os"
	"github.com/thrasher-/gocryptotrader/command"
)


func main() {
	app := command.App()
	app.Run(os.Args) // nolint: errcheck
}
