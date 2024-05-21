package rabbitmqservise

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/grpc"

	//mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	//"github.com/streadway/amqp"
	"net/url"
	// amqp "github.com/rabbitmq/amqp091-go"
	// "google.golang.org/protobuf/proto"
	//grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
)

// MqttService представляет сервис для работы с MQTT контроллером
type MqttService struct {
	conn   mqtt.Client
	server *grpc.Server
}

//func onMessageReceived(client mqtt.Client, message mqtt.Message) {
//	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
//}

// NewMqttService создает и возвращает новый экземпляр сервиса
func NewMqttService(path string, server *grpc.Server) *MqttService {

	uri, _ := url.Parse(path)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	//opts.SetOrderMatters(false)
	//opts.OnConnect = func(client mqtt.Client) {
	//	fmt.Println("Connected")
	//}
	//opts.SetDefaultPublishHandler(
	//	func(client mqtt.Client, msg mqtt.Message) {
	//		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	//	})
	//opts.OnConnectionLost = func(client mqtt.Client, err error) {
	//	fmt.Printf("Connect lost: %v", err)
	//}
	//opts.SetClientId("sub")

	c := mqtt.NewClient(opts)
	//if err != nil {
	//	log.Fatalf("NewMqttService: %s.", fmt.Errorf("%w", err))
	//}

	//c.SubscribeMultiple(nil, )
	//
	//c.StartSubscription(func(client *mqtt.MqttClient, msg mqtt.Message) {
	//	fmt.Println("Topic=", msg.Topic(), "Payload=", string(msg.Payload()))
	//})

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//c.Subscribe("cabinet/#", 0, func(client mqtt.Client, msg mqtt.Message) {
	//	result := msg.Payload()
	//	// Отправка результата в канал
	//	if result != nil {
	//		fmt.Printf("")
	//	}
	//
	//})

	//if token := c.Subscribe("cabinet/RL3H082111030142/reply/15", 0, onMessageReceived); token.Wait() && token.Error() != nil {
	//	panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	//}

	//c.Subscribe("cabinet/RL3H082111030142/replay/15", 1, func(client mqtt.Client, msg mqtt.Message) {
	//	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	//})

	return &MqttService{
		conn:   c,
		server: server,
	}
}

// PushPB публикует сообщение в формате protobuf на указанный топик и ожидает ответ от RabbitMQ
// func (s *MqttService) Subscribe(topic string, pb proto.Message) {
//
//		_ = false
//	}

func (s *MqttService) Publish(topic string, payload interface{}) mqtt.Token {
	return s.conn.Publish(topic, 0, false, payload)
}

//func (s *MqttService) Subscribe(topic string) mqtt.Token {
//	return s.conn.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
//		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
//	})
//}

func (s *MqttService) Subscribe(topic string) []byte {
	resultChannel := make(chan []byte)
	token := s.conn.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		result := msg.Payload()
		// Отправка результата в канал
		resultChannel <- result
	})

	// Ожидание завершения подписки
	if token.Wait() && token.Error() != nil {
		return nil
	}

	// Ожидание результата из колбека
	result := <-resultChannel
	return result
}
