package health

import (
	"encoding/json"
	"net/http"
)

type Status string

const (
	healthy Status = "healthy"
)

type Health struct {
	Status      Status `json:"status"`
	ServiceName string `json:"serviceName"`
}

func Handler(template Health) func(rw http.ResponseWriter, req *http.Request) {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		switch req.Method {
		case http.MethodGet:
			rw.WriteHeader(http.StatusOK)
			template.Status = healthy
			_ = json.NewEncoder(rw).Encode(template)
		case http.MethodOptions:
			rw.Header().Set("Allow", "OPTIONS, GET")
			rw.Header().Set("Cache-Control", "max-age=604800")
			rw.WriteHeader(http.StatusOK)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
