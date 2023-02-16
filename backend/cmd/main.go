package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sberapi/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config file")
}

func parseConfig() *server.Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Could not open config file", err)
	}
	config := server.NewConfig()
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Could not parse config file", err)
	}
	return config
}

func main() {
	flag.Parse()
	s := server.New(parseConfig())
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
