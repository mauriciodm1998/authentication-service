package config

import (
	"flag"
	"log"

	"github.com/notnull-co/cfg"
)

var (
	configuration config
)

type config struct {
	Token struct {
		Key string `cfg:"key"`
	} `cfg:"token"`
	Server struct {
		Port string `cfg:"port"`
	} `cfg:"server"`
	AWS struct {
		Region          string `cfg:"region"`
		AccessKeyId     string `cfg:"access_key_id"`
		SecretAccessKey string `cfg:"secret_access_key"`
		SessionToken    string `cfg:"session_token"`
	}
}

func ParseFromFlags() {
	var configDir string

	flag.StringVar(&configDir, "config-dir", "./config/", "Configuration file directory")
	flag.Parse()

	parse(configDir)
}

func parse(dirs ...string) {
	if err := cfg.Load(&configuration,
		cfg.Dirs(dirs...),
	); err != nil {
		log.Panic(err)
	}
}

func Get() config {
	return configuration
}
