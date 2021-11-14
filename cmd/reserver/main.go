package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ulstu-schedule/reserver/internal/app/reserver"
	"log"
)

const configPath = "configs/reserver.toml"

func main() {
	config := reserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	r := reserver.New(config)
	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
