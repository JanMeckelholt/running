package jung

import (
	"io"
	"net/http"
	"strings"

	"github.com/JanMeckelholt/running/http_gateway/service"
	"github.com/JanMeckelholt/running/http_gateway/service/mqtt"

	log "github.com/sirupsen/logrus"
)

func JungHandler(srv *service.Service) func(rw http.ResponseWriter, req *http.Request) {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if strings.ToLower(req.URL.Path[:5]) == service.JungRoute {
			log.Infof("jung-uri: %s", req.URL.Path)
			rw.Header().Add("Content-Type", "application/json")
			switch req.Method {
			case http.MethodPost:
				body, err := io.ReadAll(req.Body)
				if err != nil {
					log.Errorf(err.Error())
				}
				log.Infof("jung-body: %s", body)
				mqtt.Publish(srv.Clients.MqttClient, "jung", string(body))
				mqtt.Publish(srv.Clients.MqttClient, "jung", req.URL.Path)

				rw.WriteHeader(http.StatusOK)
				res, err := rw.Write(body)
				if err != nil {
					log.Errorf(err.Error())
				}
				log.Infof("Sending jung response: res: %d", res)

			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)

			}
			return
		}
	})
}
