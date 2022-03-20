package main

import (
	"flag"

	"github.com/UfaSemen/testWebSocket/pkg/server"
)

func main() {
	conf := flag.String("c", "config.toml", "configuration file path")
	flag.Parse()
	server.StartServer(*conf, server.SimpleCalculator{})
}
