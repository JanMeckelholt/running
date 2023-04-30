package logic

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/runner"
	"github.com/JanMeckelholt/running/common/grpc/strava"
)

func GetClimb(activities *database.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}

func GetWeekSummariesResponse(activities *database.ActivitiesResponse, start, end uint64) runner.WeekSummariesResponse {
	weeks := make([]*runner.WeekSummary, 0)
	startOfWeek := start
	log.Infof("GetWeekSummaries start %d -> end %d", start, end)
	for startOfWeek < end {
		log.Infof("STartOfWeek: %d", startOfWeek)
		week := GetWeek(activities, startOfWeek)
		weeks = append(weeks, &week)
		startOfWeek += 7 * 24 * 60 * 60
	}
	return runner.WeekSummariesResponse{Weeksummaries: weeks}
}

func GetWeek(activities *database.ActivitiesResponse, startOfWeek uint64) runner.WeekSummary {
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
	return week
}

func GetClimbFromStravaResponse(activities *strava.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}
