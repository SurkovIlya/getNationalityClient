package configs

import (
	"flag"
	"log"
)

type Config struct {
	Port      string
	StaticURL string
}

func New() *Config {
	cfg := Config{}
	cfg.parseFlags()

	return &cfg
}

func (c *Config) Print() {
	log.Print("\n",
		"============= CONFIG =============", "\n",
		"PORT.................: ", c.Port, "\n",
		"STATIC_URL...........: ", c.StaticURL, "\n",
	)
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.Port, "http", "8080", "")
	flag.StringVar(&c.StaticURL, "URL", "https://api.nationalize.io", "")
}
