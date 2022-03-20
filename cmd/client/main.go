package main

import (
	"flag"

	"github.com/UfaSemen/testWebSocket/pkg/client"
)

func main() {
	conf := flag.String("c", "config.toml", "configuration file path")
	flag.Parse()
	client.StartClient(*conf, client.NewRandomGenerator())
}
