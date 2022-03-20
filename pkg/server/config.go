package server

import (
	"github.com/BurntSushi/toml"
)

//Configuration structure for server, Port is the port being listened on the server
type Config struct {
	Port int
}

func readConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
