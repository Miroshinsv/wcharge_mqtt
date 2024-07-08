package mqtt

import (
	"google.golang.org/grpc"
	"log"

	rabbit "github.com/Miroshinsv/wcharge_mqtt/pkg/rabbitmq_service"
)

type MqttController struct {
	Rabbit *rabbit.MqttService
	/*logger logger.Interface*/
	seqs map[string]int // нумерация пакетов по топикам (для rl_seq)
}

func NewMqttController(url string /*l logger.Interface, */, server *grpc.Server) *MqttController {
	return &MqttController{
		Rabbit: rabbit.NewMqttService(url, server),
		seqs:   map[string]int{},
		/*logger: l,*/
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
