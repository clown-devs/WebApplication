package server

type Config struct {
	BindAddr string `json:"bindAddr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
	}
}
