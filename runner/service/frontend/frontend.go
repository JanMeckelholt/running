package frontend

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/JanMeckelholt/running/common/grpc/runner"

	log "github.com/sirupsen/logrus"
)

type TemplateData struct {
	RunningData *runner.WeekSummary
	Week        int64
	Year        int64
}

var (
	tmpl *template.Template
	err  error
)
var templPath = "./runner/service/frontend/html-templates/"
var templFile = "runner.html"
var funcMap = template.FuncMap{
	"decWeek": func(year, week int64) int64 {
		if week <= 1 {
			return 1
		}
		return week - 1
	},
	"incWeek": func(year, week int64) int64 {
		isoYear, isoWeek := time.Now().ISOWeek()
		if year == int64(isoYear) && week >= int64(isoWeek) {
			return int64(isoWeek)
		}
		return week + 1
	},
	"decYear": func(year, week int64) int64 {
		if year <= 1900 {
			return 1900
		}
		return year - 1
	},
	"incYear": func(year, week int64) int64 {
		isoYear, _ := time.Now().ISOWeek()
		if year >= int64(isoYear) {
			return int64(isoYear)
		}
		if year <= 1900 {
			return 1900
		}
		return year + 1
	},
	"distToStr": func(distance uint64) string {
		return fmt.Sprintf("%.2f km", float64(distance)/1000)
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
		var weekInt, yearInt int
		weekStr := r.URL.Query().Get("week")
		yearStr := r.URL.Query().Get("year")
		log.Infof("Query: %s", r.URL)
		weekInt, err := strconv.Atoi(weekStr)
		if err != nil {
			yearInt, weekInt = time.Now().ISOWeek()
		} else {
			yearInt, err = strconv.Atoi(yearStr)
			if err != nil {
				yearInt, weekInt = time.Now().ISOWeek()
			}
		}

		log.Infof("Year: %d, Week: %d", yearInt, weekInt)
		data, err := rs.GetWeekSummary(context.Background(), &runner.WeekSummaryRequest{ClientId: "77376", Week: uint64(weekInt), Year: uint64(yearInt)})
		if err != nil {
			log.Errorf("Error getting Weeksummary: %s", err.Error())
		}
		err = tmpl.Execute(w, TemplateData{RunningData: data, Week: int64(weekInt), Year: int64(yearInt)})
		if err != nil {
			log.Errorf("Could not parse template: %s", err.Error())
		}
	}

}
