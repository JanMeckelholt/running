package logic

import (
	"time"

	"github.com/JanMeckelholt/running/common/grpc/database"
	"github.com/JanMeckelholt/running/common/grpc/strava"
)

const startOFWeek = time.Monday

func GetStartOfWeek() uint64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	daysSinceMonday := time.Duration(roundDay.Weekday() - startOFWeek)
	return uint64(roundDay.Add(-daysSinceMonday * 24 * time.Hour).Unix())
}

func GetClimb(activities *database.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}

func GetClimbFromStravaResponse(activities *strava.ActivitiesResponse) uint64 {
	var totalClimb uint64
	for _, activity := range activities.GetActivities() {
		totalClimb += uint64(activity.GetTotalElevationGain())
	}
	return totalClimb
}
