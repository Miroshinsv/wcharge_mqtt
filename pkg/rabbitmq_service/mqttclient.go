package rabbitmqservise

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_mqtt/config"
	grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"strconv"
	"time"

	//"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"strings"

	//mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	//"github.com/streadway/amqp"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/url"
	// "google.golang.org/protobuf/proto"
	//grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
)

// MqttService представляет сервис для работы с MQTT контроллером
type MqttService struct {
	conn   mqtt.Client
	server *grpc.Server
	rabbit *amqp.Connection
}

//func onMessageReceived(client mqtt.Client, message mqtt.Message) {
//	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
//}

// NewMqttService создает и возвращает новый экземпляр сервиса
func NewMqttService(path string, server *grpc.Server) *MqttService {

	cfg, _ := config.NewConfig()
	cc, _ := amqp.Dial(cfg.Rabbit.URL)
	defer cc.Close()

	uri, _ := url.Parse(path)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	s := MqttService{
		conn:   c,
		server: server,
		rabbit: cc,
	}

	c.Subscribe("cabinet/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		//result := msg.Payload()
		//// Отправка результата в канал
		//if result != nil {
		//	fmt.Printf("")
		//}

		topicAr := strings.Split(msg.Topic(), "/")

		//if topicAr[2] == "report" && topicAr[3] == "22" {
		//	ms := &grpc_v1.PBReturnReportMsg{}
		//	err := proto.Unmarshal(msg.Payload(), ms)
		//	if err != nil {
		//		log.Printf("MQTT. Ошибка коннекта станции: %s", err)
		//	}
		//}
		//if else topicAr[2] == "report" && topicAr[3] == "10" {
		//
		//}

		switch {
		case topicAr[2] == "report" && topicAr[3] == "10":
			ms := &grpc_v1.RequestReportCabinetLogin{}
			err := proto.Unmarshal(msg.Payload(), ms)
			if err != nil {
				log.Printf("MQTT. Ошибка коннекта станции: %s", err)
			}

			time.Sleep(60 * time.Second) // TODO

			s.sendNewStationConnect(ms, topicAr[1])
		case topicAr[2] == "report" && topicAr[3] == "22":
			ms := &grpc_v1.PBReturnReportMsg{}
			err := proto.Unmarshal(msg.Payload(), ms)
			if err != nil {
				log.Printf("MQTT. Ошибка коннекта станции: %s", err)
			}
		}

		//if msg.Topic() == "cabinet/RL3H082111030142/report/22" {
		//	//result := msg.Payload()
		//	ms := &grpc_v1.RequestReturnPowerBank{}
		//	err := proto.Unmarshal(msg.Payload(), ms)
		//	if err != nil {
		//		fmt.Printf("")
		//	}
		//}
		//
		//if msg.Topic() != "cabinet/RL3H082111030142/report/22" {
		//	if msg.Topic() == "cabinet/RL3H082111030142/report/10" {
		//		//result := msg.Payload()
		//		ms := &grpc_v1.RequestReportCabinetLogin{}
		//		err := proto.Unmarshal(msg.Payload(), ms)
		//		if err != nil {
		//			fmt.Printf("")
		//		}
		//	}
		//}

	})

	//if token := c.SubscribeMqtt("cabinet/RL3H082111030142/reply/15", 0, onMessageReceived); token.Wait() && token.Error() != nil {
	//	panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	//}
	//c.SubscribeMqtt("cabinet/RL3H082111030142/replay/15", 1, func(client mqtt.Client, msg mqtt.Message) {
	//	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	//})

	s.test()

	//messageBytes, _ := proto.Marshal(&grpc_v1.RequestInventory{})
	//s.PublishMqtt("cabinet/RL3H082111030142/cmd/13", messageBytes)
	//
	//res := s.SubscribeMqtt("cabinet/RL3H082111030142/reply/13")
	//msg := &grpc_v1.ResponseInventory{}
	//err := proto.Unmarshal(res, msg)
	//
	//if err != nil {
	//	log.Printf("MQTT. Ошибка коннекта станции: %s", err)
	//}

	//
	return &s
}

func (s *MqttService) sendNewStationConnect(ms *grpc_v1.RequestReportCabinetLogin, stationName string) {
	topicStart := "cabinet/" + stationName + "/"

	topicCmdGetAllPowerbanks := topicStart + "cmd/13"
	topicReplyGetAllPowerbanks := topicStart + "reply/13"

	messageBytes, _ := proto.Marshal(&grpc_v1.RequestInventory{})

	//time.Sleep(60 * time.Second)

	s.PublishMqtt(topicCmdGetAllPowerbanks, messageBytes)
	res := s.SubscribeMqtt(topicReplyGetAllPowerbanks)
	msg := &grpc_v1.ResponseInventory{}
	err := proto.Unmarshal(res, msg)

	if err != nil {
		log.Printf("MQTT. Ошибка коннекта станции: %s", err)
	}

	type Powerbank struct {
		Position     int
		SerialNumber string
		Capacity     int
		Used         int
	}

	type FullStation struct {
		SerialNumber string
		Capacity     int
		Powerbanks   []Powerbank
	}

	station := FullStation{
		SerialNumber: stationName,
		Capacity:     int(ms.RlCount),
	}

	for i, p := range msg.Slot {
		fmt.Println(i, p)
		if p.RlPbid != 0 {
			powerbank := Powerbank{
				Position:     int(p.RlSlot),
				SerialNumber: strconv.Itoa(int(p.RlPbid)),
			}
			station.Powerbanks = append(station.Powerbanks, powerbank)
		}
	}

	log.Printf("Добавлена новая станция: %s", stationName)

	s.PublishRabbit("mqtt_add_station", station)
}

// TODO \/
func (s *MqttService) test() {
	topicAr := strings.Split("cabinet/RL3H082111030142/report/10", "/")
	ms := &grpc_v1.RequestReportCabinetLogin{
		RlCount:       8, // Capacity
		RlNetmode:     1,
		RlConn:        3,
		RlCsq:         62,
		RlRsrp:        171,
		RlSinr:        108,
		RlWifi:        0,
		RlCommsoftver: "RL1.M6.08.04",
		RlCommhardver: "ff",
		RlIccid:       "89701010050648321412",
		RlSeq:         1,
	}

	if topicAr != nil {
	}
	if ms != nil {
	}

	// 						   SerialNumber
	topicStart := "cabinet/" + topicAr[1] + "/"

	topicCmdGetAllPowerbanks := topicStart + "cmd/13"
	topicReplyGetAllPowerbanks := topicStart + "reply/13"

	messageBytes, _ := proto.Marshal(&grpc_v1.RequestInventory{})

	s.PublishMqtt(topicCmdGetAllPowerbanks, messageBytes)
	res := s.SubscribeMqtt(topicReplyGetAllPowerbanks)
	msg := &grpc_v1.ResponseInventory{}
	err := proto.Unmarshal(res, msg)

	if err != nil {

	}

	type Powerbank struct {
		Position     int
		SerialNumber string
		Capacity     int
		Used         int
	}

	type FullStation struct {
		SerialNumber string
		Capacity     int
		Powerbanks   []Powerbank
	}

	station := FullStation{
		SerialNumber: topicAr[1],
		Capacity:     int(ms.RlCount),
	}

	for i, p := range msg.Slot {
		fmt.Println(i, p)
		if p.RlPbid != 0 {
			powerbank := Powerbank{
				Position:     int(p.RlSlot),
				SerialNumber: strconv.Itoa(int(p.RlPbid)),
			}
			station.Powerbanks = append(station.Powerbanks, powerbank)
		}
	}

	log.Printf("Добавлена новая станция: %s", topicAr[1])

	//if station.Capacity != 0 {
	//}

	//type Data struct {
	//	T int
	//	D string
	//}
	//
	//data := Data{
	//	T: 1,
	//	D: "adsadasd",
	//}

	//res = s.SubscribeMqtt("test/mqtt")
	s.PublishRabbit("mqtt_add_station", station)
	//
	//if res != nil {
	//
	//}
	//ch, _ := s.rabbit.Channel()
	//defer ch.Close()
	//ch.QueueDeclare(
	//	"test/rabbit", // name
	//	false,         // durable
	//	false,         // delete when unused
	//	true,          // exclusive
	//	false,         // noWait
	//	nil,           // arguments
	//)
	//
	//msgs, err := ch.Consume(
	//	"test/rabbit", // queue
	//	"",            // consumer
	//	true,          // auto-ack
	//	false,         // exclusive
	//	false,         // no-local
	//	false,         // no-wait
	//	nil,           // args
	//)
	//
	//body, _ := json.Marshal(data)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//ch.PublishWithContext(
	//	ctx,
	//	"",
	//	"test/rabbit",
	//	false,
	//	false,
	//	amqp.Publishing{
	//		ContentType: "application/json",
	//		Body:        body,
	//	},
	//)
	//
	//var mssg proto.Message
	//for d := range msgs {
	//	err := json.Unmarshal(d.Body, mssg)
	//	if err != nil {
	//		continue
	//	}
	//}
}

// TODO /\

func (s *MqttService) PublishMqtt(topic string, payload interface{}) mqtt.Token {
	return s.conn.Publish(topic, 0, false, payload)
}

func (s *MqttService) SubscribeMqtt(topic string) []byte {
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

func (s *MqttService) PublishRabbit(topic string, payload interface{}) bool {
	ch, _ := s.rabbit.Channel()
	defer ch.Close()
	err := ch.ExchangeDeclare(
		topic,   // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	//messageBytes, err := proto.Marshal(payload)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//type Data struct {
	//	T int
	//	D string
	//}
	//
	//data := Data{
	//	T: 1,
	//	D: "adsadasd",
	//}

	body, _ := json.Marshal(payload)
	os.Stdout.Write(body)
	err = ch.PublishWithContext(
		ctx,
		topic,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        body,
		},
	)

	if err != nil {
		return false
	}

	return true
}

func (s *MqttService) SubscribeRabbit(topic string) proto.Message {
	ch, _ := s.rabbit.Channel()
	ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	msgs, _ := ch.Consume(
		topic, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	//json.Un

	//var msgs
	var msg proto.Message
	for d := range msgs {
		err := json.Unmarshal(d.Body, msg)
		if err != nil {
			continue
		}
	}

	return msg
}
