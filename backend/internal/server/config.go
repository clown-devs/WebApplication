package server

import "github.com/sirupsen/logrus"

type Config struct {
	BindAddr string `json:"bindAddr"`
	LogLevel string `json:"logLevel"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: logrus.DebugLevel.String(),
	}
}
