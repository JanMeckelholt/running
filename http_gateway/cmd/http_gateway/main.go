package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/JanMeckelholt/running/common/commonconf"
	"github.com/JanMeckelholt/running/common/dependencies"
	"github.com/JanMeckelholt/running/common/utils"
	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/auth"
	"github.com/JanMeckelholt/running/http_gateway/service/config"
	"github.com/JanMeckelholt/running/http_gateway/service/mqtt"

	"github.com/JanMeckelholt/running/http_gateway/service/mux"
	regexphandler "github.com/JanMeckelholt/running/http_gateway/service/regexpHandler"
	"github.com/JanMeckelholt/running/http_gateway/service/server"
	"github.com/joho/godotenv"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

func main() {
	commonConf := commonconf.GPGConf{}
	err := env.Parse(&commonConf)
	if err != nil {
		return
	}
	err = utils.DecryptPGP("./common/.env.docker.running_app.secret.asc", "./secret/env/.env.docker.running_app.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt common-env: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/env/.env.docker.secret.asc", "./secret/env/.env.docker.secret", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt env: %s", err.Error())
		return
	}
	err = utils.DecryptPGP("./volumes-data/certs/http_gateway-server-key.pem.asc", "./secret/certs/http_gateway-server-key.pem", commonConf.GPGPrivateKey)
	if err != nil {
		log.Errorf("Could not decrypt certs: %s", err.Error())
		return
	}
	err = godotenv.Load("./volumes-data/env/.env.docker", "./secret/env/.env.docker.secret", "./common/.env.docker.running_app", "./secret/env/.env.docker.running_app.secret")
	if err != nil {
		log.Errorf("Could not load env: %s", err.Error())
		return
	}
	srv := &service.Service{}
	err = env.Parse(&srv.Config)
	if err != nil {
		log.Errorf("Could not load serviceConfig: %s", err.Error())
	}
	log.Infof("RunningAppPort: %d", srv.Config.RunningAppPort)
	if len(srv.Config.JWTKey) < 8 {
		log.Fatal("JWT not valid")
	}
	auth.JwtKey = []byte(srv.Config.JWTKey)

	err = auth.UpsertRunningClients(auth.RunningAppHttpClient, srv.Config.RunningAppPassword)
	if err != nil {
		log.Errorf("Could not setup password for runningApp-Client %s", auth.RunningAppHttpClient)
	}
	err = auth.UpsertHttpClients(auth.MasterHttpClient, srv.Config.MasterPassword)
	if err != nil {
		log.Errorf("Could not setup password for HTTP-Client %s", auth.MasterHttpClient)
	}

	err = srv.Clients.Dial(srv.Config)
	if err != nil {
		log.Errorf("could not dial clients! %s", err.Error())
	}
	rs, err := server.NewHttpGatewayServer(srv.Clients)
	if err != nil {
		log.Errorf("could create gateway server! %s", err.Error())
	}

	apiHandler := rHandler(rs, srv)
	go mqtt.ServeMqtt(srv)
	serveTLS(apiHandler, dependencies.Configs["http_gateway-API"].Port)

}

func serveTLS(handler http.Handler, addr int) {
	tlsServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", addr),
		Handler: handler,
	}
	lis, err := net.Listen("tcp", tlsServer.Addr)
	if err != nil {
		return
	}

	teardown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdownErr := tlsServer.Shutdown(ctx)
		if shutdownErr != nil {
			log.Fatal("TLS-server: Shutdown Error!")
		}
	}

	log.Infof("Listening on :%d", addr)
	serveErr := tlsServer.ServeTLS(lis, "volumes-data/certs/http_gateway-server-cert.pem", "secret/certs/http_gateway-server-key.pem")
	defer func() {
		teardown()
	}()
	if serveErr != nil {
		log.Fatalf("HTTP-Gateway-Server: Serving Error: %s", serveErr.Error())
	}
}

func rHandler(rs *server.HttpGatewayServer, srv *service.Service) http.Handler {
	r := regexphandler.RegexpHandler{}

	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", service.LoginRoute)), mux.Handler(service.LoginRoute, rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s.*$", service.JungRoute)), mux.Handler(service.JungRoute, rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/health")), mux.Handler("/health", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/athlete")), mux.Handler("/athlete", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/activities")), mux.Handler("/activities", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/athlete/create")), mux.Handler("/athlete/create", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/client/create")), mux.Handler("/client/create", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/weeksummary")), mux.Handler("/weeksummary", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/weeksummaries")), mux.Handler("/weeksummaries", rs, srv))
	r.HandleFunc(regexp.MustCompile(fmt.Sprintf("^%s$", config.ApiPrefix+config.RunPrefix+"/activitiesToDB")), mux.Handler("/activitiesToDB", rs, srv))

	rWithAuth := server.AuthMiddleware(r)

	return rWithAuth
}
