package mqtt

import (
	"log"

	rabbit "github.com/Miroshinsv/wcharge_mqtt/pkg/rabbitmq_servise"
)

type MqttController struct {
	rabbit *rabbit.MqttService
	seqs   map[string]int // нумерация пакетов по топикам (для rl_seq)
}

func NewMqttController(url string) *MqttController {
	return &MqttController{
		rabbit: rabbit.NewMqttService(url),
		seqs:   map[string]int{},
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
