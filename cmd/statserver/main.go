package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/amartery/statSaver/internal/app/statserver"
)

var (
	configPath string = "./configs/statserver.toml"
)

func main() {
	config := statserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := statserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
