package main

import (
	"github.com/ulstu-schedule/reserver/internal/app"
)

const configsPath = "configs"

func main() {
	app.Run(configsPath)
}
