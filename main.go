package main

import (
	"e-commerce-system-using-go/config"
	"e-commerce-system-using-go/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Config file is not loaded properly: %v\n", err);
		return
	}

	api.StartServer(cfg)
}