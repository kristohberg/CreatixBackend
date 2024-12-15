package main

import (
	"log"

	"github.com/kristohberg/CreatixBackend/api"
	"github.com/kristohberg/CreatixBackend/config"
	"github.com/kristohberg/CreatixBackend/logging"
)

//main

func main() {
	standardLogger := logging.NewLogger()
	// Set up config
	cfg, err := config.SetUpConfig()
	if err != nil {
		log.Fatalf("Not able to set config: %s ", err.Error())
		standardLogger.Misconfigured("Configuration is misconfigured", err)
		return
	}

	a, err := api.New(cfg)
	if err != nil {
		standardLogger.Error(err)
		log.Fatal(err.Error())
		return
	}

	a.Run()
}
