package clients

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	certhandling "github.com/JanMeckelholt/running/common/cert-handling"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/http_gateway/service/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Clients struct {
	RunnerClient runner.RunnerClient
	MqttClient   mqtt.Client
}

func (c *Clients) Dial(config config.ServiceConfig) error {
	conn, err := dial(config.RunnerName, "runner")
	if err != nil {
		return err
	}
	c.RunnerClient = runner.NewRunnerClient(conn)

	return nil
}

func dial(serviceName, address string) (*grpc.ClientConn, error) {
	tlsCredentials, err := certhandling.LoadTLSClientCredentials("volumes-data/certs/ca-cert.pem")
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	log.Infof("Dialing: %s:%d", address, dependencies.Configs[serviceName].Port)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", address, dependencies.Configs[serviceName].Port), grpc.WithTransportCredentials(tlsCredentials))
	log.Infof("Dialed  %s", conn.Target())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
