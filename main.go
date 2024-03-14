package main

import (
	"github.com/marcelofpfelix/nika/core"
)

func main() {
	config := core.Config{}

	config = core.ReadConfig()
	core.ConfigCli(config)
}
