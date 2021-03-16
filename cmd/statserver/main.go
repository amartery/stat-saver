package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/amartery/statSaver/internal/app/delivery/http"
	"github.com/amartery/statSaver/internal/app/repository/postgresDB"
	"github.com/amartery/statSaver/internal/app/usecase"
	"github.com/amartery/statSaver/internal/pkg/utility"
)

var (
	configPath string = "./configs/statserver.toml"
)

func main() {
	config := http.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	postgresCon, err := utility.CreatePostgresConnection(config.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	statRep := postgresDB.NewStatRepository(postgresCon)
	statUsecase := usecase.NewStatUsecase(statRep)

	s := http.New(config, statUsecase)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
