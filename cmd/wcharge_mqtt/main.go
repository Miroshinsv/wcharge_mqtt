package main

import (
	"log"

	"github.com/Miroshinsv/wcharge_mqtt/config"
	"github.com/Miroshinsv/wcharge_mqtt/internal/app/wcharge_mqtt"
)

// Run the wcharge_mqtt app
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	wcharge_mqtt.Run(cfg)
}
