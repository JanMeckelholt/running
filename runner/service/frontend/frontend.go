package frontend

import (
	"context"
	"html/template"
	"net/http"

	"github.com/JanMeckelholt/running/common/grpc/runner"

	log "github.com/sirupsen/logrus"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
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

func FrontEnd(week int64, rs runner.RunnerServer) http.HandlerFunc {
	data, err := rs.GetWeekSummary(context.Background(), &runner.WeekSummaryRequest{ClientId: "77376", Week: week})
	if err != nil {
		log.Errorf("Error getting Weeksummary: %s", err.Error())
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, data)
	}

}
