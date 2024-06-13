package main

import (
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
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

	graylogAddr := cfg.Graylog.URL
	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewUDPWriter: %s", err)
	}

	log.SetOutput(gelfWriter)
	log.Println("MQTT. Это информационное сообщение")

	wcharge_mqtt.Run(cfg)
}
