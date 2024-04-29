package main

import (
	"github.com/marcelofpfelix/boa/core"
)

func main() {
	config := core.Config{}

	config = core.ReadConfig()
	core.ConfigCli(config)
}
