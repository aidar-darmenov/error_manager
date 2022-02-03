package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	HttpPort     int
	ListenerHost string
	GrpcPort     int `json:"GrpcPort"`
}

//NewConfiguration read file, return configuration
func NewConfiguration(path string) *Configuration {
	var configuration Configuration
	configuration.InitConfigParams(path)
	return &configuration
}

func (c *Configuration) InitConfigParams(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Configuration) Params() *Configuration {
	return c
}
