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

func init() {
	tmpl, err = template.ParseFiles("./runner/service/frontend/html-templates/runner.html")
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
	}

}
