package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/MuhahaSam/golangPractice/internal/app"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./note_config.json", "Path to configuration file")
	time.Local = time.UTC
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("Can't create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Can't run app: %s", err.Error())
	}
}
