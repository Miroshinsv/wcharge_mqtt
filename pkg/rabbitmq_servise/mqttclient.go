package rabbitmqservise

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

// MqttService представляет сервис для работы с MQTT контроллером
type MqttService struct {
	conn *amqp.Connection
}

// NewMqttService создает и возвращает новый экземпляр сервиса
func NewMqttService(url string) *MqttService {
	c, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("NewMqttService: %s.", fmt.Errorf("%w", err))
	}

	return &MqttService{
		conn: c,
	}
}

// PushPB публикует сообщение в формате protobuf на указанный топик и ожидает ответ от RabbitMQ
func (s *MqttService) Subscribe(topic string, pb proto.Message) {

}

func (s *MqttService) Channel() (*amqp.Channel, error) {
	return s.conn.Channel()
}
