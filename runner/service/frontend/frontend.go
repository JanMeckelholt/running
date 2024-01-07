package frontend

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JanMeckelholt/running/common/grpc/runner"

	log "github.com/sirupsen/logrus"
)

type TemplateData struct {
	RunningData *runner.WeekSummary
	Week        int
	Year        int
}

type yearWeek struct {
	Year int
	Week int
}

var (
	tmpl *template.Template
	err  error
)
var templPath = "./runner/service/frontend/html-templates/"
var templFile = "runner.html"
var funcMap = template.FuncMap{
	"decWeek": func(year, week int) yearWeek {
		t := yearWeekToDate(year, week)
		t = t.AddDate(0, 0, -7)
		yW := sanityCheckDate(t)
		if yW != nil {
			return *yW
		}
		y, w := t.ISOWeek()
		return yearWeek{y, w}
	},

	"incWeek": func(year, week int) yearWeek {
		t := yearWeekToDate(year, week)
		t = t.AddDate(0, 0, 7)
		yW := sanityCheckDate(t)
		if yW != nil {
			return *yW
		}
		y, w := t.ISOWeek()
		return yearWeek{y, w}
	},

	"decYear": func(year, week int) yearWeek {
		t := time.Date(year-1, time.January, 4, 0, 0, 0, 0, time.Local)
		yW := sanityCheckDate(t)
		if yW != nil {
			return *yW
		}
		y, w := t.ISOWeek()
		return yearWeek{y, w}
	},
	"incYear": func(year, week int) yearWeek {
		t := time.Date(year+1, time.January, 4, 0, 0, 0, 0, time.Local)
		yW := sanityCheckDate(t)
		if yW != nil {
			return *yW
		}
		y, w := t.ISOWeek()
		return yearWeek{y, w}
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
		if strings.Contains(r.URL.String(), ".css") {
			log.Errorf("This should not be handled in Frontend Handler: %s", r.URL.String())
		}
		weekInt, err := strconv.Atoi(weekStr)
		if err != nil {
			log.Errorf("error week str: %s", err.Error())
			yearInt, weekInt = time.Now().ISOWeek()
		} else {
			yearInt, err = strconv.Atoi(yearStr)
			if err != nil {
				log.Errorf("error year str: %s", err.Error())
				yearInt, weekInt = time.Now().ISOWeek()
			}
		}

		log.Infof("Year: %d, Week: %d", yearInt, weekInt)
		data, err := rs.GetWeekSummary(context.Background(), &runner.WeekSummaryRequest{ClientId: "77376", Week: uint64(weekInt), Year: uint64(yearInt)})
		if err != nil {
			log.Errorf("Error getting Weeksummary: %s", err.Error())
		}
		err = tmpl.Execute(w, TemplateData{RunningData: data, Week: weekInt, Year: yearInt})
		if err != nil {
			log.Errorf("Could not parse template: %s", err.Error())
		}
	}

}

func CSSHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("../../frontend/html-templates"))
	}
}

func yearWeekToDate(year, week int) time.Time {
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.Local)
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)
	return t
}

func sanityCheckDate(t time.Time) *yearWeek {
	if t.After(time.Now()) {
		log.Infof("time %s after current date, defaulting to current date", t.String())
		y, w := time.Now().ISOWeek()
		return &yearWeek{y, w}
	}
	if t.Before(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)) {
		log.Infof("time %s before 1900, defaulting to first week of 1900", t.String())
		return &yearWeek{1900, 1}
	}
	return nil
}
