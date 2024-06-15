package configs

import (
	"flag"
	"log"
)

type Config struct {
	Server struct {
		Port      string
		StaticURL string
	}
	NationalData struct {
		Host string
	}
	Cache struct {
		TimeToLifeCache int
		Count           int
	}
}

func New() *Config {
	cfg := Config{}
	cfg.parseFlags()

	return &cfg
}

func (c *Config) Print() {
	log.Print("\n",
		"============= CONFIG =============", "\n",
		"PORT.................: ", c.Server.Port, "\n",
		"STATIC_URL...........: ", c.Server.StaticURL, "\n",

		//NationalData
		"HOST...........: ", c.NationalData.Host, "\n",

		//Cahce
		"TTL_CACHE...........: ", c.Cache.TimeToLifeCache, "\n",
		"COUNT...........: ", c.Cache.Count, "\n",
	)
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.Server.Port, "http", "8080", "")
	flag.StringVar(&c.Server.StaticURL, "URL", "localhost", "")
	flag.StringVar(&c.NationalData.Host, "Host", "https://api.nationalize.io", "")
	flag.IntVar(&c.Cache.TimeToLifeCache, "ttlCache", 5, "")
	flag.IntVar(&c.Cache.Count, "Count", 1000, "")
}
