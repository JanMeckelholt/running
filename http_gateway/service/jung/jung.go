package jung

import (
	"io"
	"net/http"
	"strings"

	"github.com/JanMeckelholt/running/http_gateway/service"
	log "github.com/sirupsen/logrus"
)

func JungHandler() func(rw http.ResponseWriter, req *http.Request) {
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
