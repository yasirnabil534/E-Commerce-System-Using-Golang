package main

import (
	"e-commerce-system-using-golang/config"
	"e-commerce-system-using-golang/internal/api"
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