package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nvlhnn/gommerce/internal/api"
	"github.com/nvlhnn/gommerce/internal/config"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server := api.NewServerHTTP(config)
	server.Start()
	
}