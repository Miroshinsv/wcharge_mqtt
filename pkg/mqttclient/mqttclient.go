package mqttclient

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client MQTT.Client
}

func NewMqttClient(broker_url string, client_id string) *MqttClient {
	opts := MQTT.NewClientOptions().AddBroker(broker_url)
	opts.SetClientID(client_id)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	return &MqttClient{client: client}
}

func (mc *MqttClient) Subscribe(topic string) {
	if token := mc.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func (mc *MqttClient) Unsubscribe(topic string) {
	if token := mc.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Printf("Unsubscribed from topic: %s\n", topic)
}

func (mc *MqttClient) Disconnect() {
	mc.client.Disconnect(250) // Wait 250 milliseconds before disconnecting
}
