package mqtt

import (
	"fmt"
	"time"

	"github.com/JanMeckelholt/running/http_gateway/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func ServeMqtt(srv *service.Service) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("mosquitto:1883")
	//opts.AddBroker("192.168.178.61:1883")
	opts.SetClientID("go_mqtt_client")
	//opts.SetUsername(srv.Config.MqttUserName)
	//opts.SetPassword(srv.Config.MasterPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	srv.Clients.MqttClient = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Errorf("user: __%s__, password: __%s__", srv.Config.MqttUserName, srv.Config.MqttPassword)
		panic(token.Error())
	}

}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func Publish(client mqtt.Client, topic string, message string) {
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	time.Sleep(time.Second)
}
