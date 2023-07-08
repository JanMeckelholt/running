package frontend

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/JanMeckelholt/running/common/grpc/runner"

	log "github.com/sirupsen/logrus"
)

type TemplateData struct {
	RunningData *runner.WeekSummary
	Week        int64
}

var (
	tmpl *template.Template
	err  error
)
var templPath = "./runner/service/frontend/html-templates/"
var templFile = "runner.html"
var funcMap = template.FuncMap{
	"dec": func(Week int64) int64 {
		return Week - 1
	},
	"inc": func(Week int64) int64 {
		if Week >= 0 {
			return 0
		}
		return Week + 1
	},
}

func init() {
	tmpl, err = template.New(templFile).Funcs(funcMap).ParseFiles(templPath + templFile)
	if err != nil {
		panic(err)
	}
}

func FrontEnd(rs runner.RunnerServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		weekStr := r.URL.Query().Get("week")
		weekInt, err := strconv.Atoi(weekStr)
		if err != nil {
			weekInt = 0
		}
		data, err := rs.GetWeekSummary(context.Background(), &runner.WeekSummaryRequest{ClientId: "77376", Week: int64(weekInt)})
		if err != nil {
			log.Errorf("Error getting Weeksummary: %s", err.Error())
		}
		err = tmpl.Execute(w, TemplateData{RunningData: data, Week: int64(weekInt)})
		if err != nil {
			log.Errorf("Could not parse template: %s", err.Error())
		}
	}

}
