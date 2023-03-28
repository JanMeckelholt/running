package logic

import (
	"time"

	"github.com/JanMeckelholt/running/common/grpc/strava"
)

const startOFWeek = time.Monday

func GetStartOfWeek() int64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	daysSinceMonday := time.Duration(roundDay.Weekday() - startOFWeek)
	return roundDay.Add(-daysSinceMonday * 24 * time.Hour).Unix()
}

func GetClimb(activities *strava.ActivitiesResponse) int64 {
	var totalClimb int64
	for _, activity := range activities.GetActivities() {
		totalClimb += int64(activity.GetTotalElevationGain())
	}
	return totalClimb
}
