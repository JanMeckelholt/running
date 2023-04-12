package logic

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
)

const startOFWeek = time.Monday

func GetStartOfFirstWeek(weeksAgo uint64) uint64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	daysSinceMonday := time.Duration(roundDay.Weekday() - startOFWeek)
	return uint64(roundDay.Add(-daysSinceMonday*24*time.Hour - time.Duration(weeksAgo*7*24*uint64(time.Hour))).Unix())
}

func GetClimb(activities *database.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}

func GetWeekSummarryResponse(activities *database.ActivitiesResponse, start uint64) runner.WeekSummariesResponse {

	weeks := make([]*runner.WeekSummary, 0)
	startOfWeek := start
	now := uint64(time.Now().Unix())
	log.Infof("now: %d", now)
	for startOfWeek < now {
		log.Infof("STartOfWeek: %d", startOfWeek)
		var (
			totalClimb, NumberOfRuns, numberOfOthers, distance, elapsedTime uint64
		)
		for _, activity := range activities.GetActivities() {
			if activity.GetStartDateUnix() < startOfWeek || activity.GetStartDateUnix() > startOfWeek+(7*24*60*60) {
				continue
			}
			totalClimb += uint64(activity.GetTotalElevationGain())
			elapsedTime += uint64(activity.GetElapsedTime())
			distance += uint64(activity.GetDistance())
			if activity.GetType() == "Lauf" || activity.GetType() == "Run" {
				NumberOfRuns++
			} else {
				numberOfOthers++
			}
		}
		week := runner.WeekSummary{
			Distance:       distance,
			TimeUnix:       elapsedTime,
			NumberOfRuns:   NumberOfRuns,
			Climb:          totalClimb,
			NumberOfOthers: numberOfOthers,
		}
		t := time.Duration(week.TimeUnix * uint64(time.Second))
		timeStr := t.String()
		week.TimeStr = timeStr
		weeks = append(weeks, &week)
		startOfWeek += 7 * 24 * 60 * 60
	}
	return runner.WeekSummariesResponse{Weeksummaries: weeks}
}

func GetClimbFromStravaResponse(activities *strava.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}
