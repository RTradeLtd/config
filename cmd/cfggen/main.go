package main

import (
	"flag"

	"github.com/RTradeLtd/config"
)

func main() {
	var cfgPath = flag.String("cfgpath", "config.json", "path to generate config to")
	flag.Parse()
	config.GenerateConfig(*cfgPath)
}
