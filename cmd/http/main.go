package main

import (
	"log"
	"refactoring/config"
	"refactoring/internl/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(cfg)
}
