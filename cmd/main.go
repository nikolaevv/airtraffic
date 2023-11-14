package main

import (
	"flag"
	"github.com/nikolaevv/airtraffic/internal/app"
	"log"
)

var (
	cfgPath = flag.String("cfg", "./config/config.json", "path to config file")
)

func main() {
	s, err := app.New(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
