package mqtt

import (
	"log"

	"github.com/Miroshinsv/wcharge_mqtt/pkg/logger"
	rabbit "github.com/Miroshinsv/wcharge_mqtt/pkg/rabbitmq_servise"
)

type MqttController struct {
	rabbit *rabbit.MqttService
	logger logger.Interface
	seqs   map[string]int // нумерация пакетов по топикам (для rl_seq)
}

func NewMqttController(url string, l logger.Interface) *MqttController {
	return &MqttController{
		rabbit: rabbit.NewMqttService(url),
		seqs:   map[string]int{},
		logger: l,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
