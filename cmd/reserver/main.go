package main

import (
	"github.com/ulstu-schedule/reserver/internal/app/config"
	"github.com/ulstu-schedule/reserver/internal/app/reserver"
	"log"
)

const configsPath = "configs"

func main() {
	cfg, err := config.Init(configsPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = reserver.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
