package client

import (
	"github.com/BurntSushi/toml"
)

/*
Configuration structure for client
ServerAddress is the network address of the server
ServerPort is the port of the server
OutputPath is the local path for the output files
ConnectionNum is the number of connections to server
MessageNumPerCon is the number of messages being sent to the server in one connection
RateLimit is the limit of the rate of message sending to the server in milliseconds
*/
type Config struct {
	ServerAddress    string
	ServerPort       int
	OutputPath       string
	ConnectionNum    int
	MessageNumPerCon int
	RateLimit        int
}

func readConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
