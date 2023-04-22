package utils

import "time"

const startOFWeek = time.Monday

func GetStartOfFirstWeek(weeksAgo uint64) uint64 {
	roundDay := time.Now().Round(24 * time.Hour)
	if roundDay.Weekday() != time.Now().Weekday() {
		roundDay = roundDay.Add(-24 * time.Hour)
	}
	daysSinceMonday := time.Duration(roundDay.Weekday() - startOFWeek)
	return uint64(roundDay.Add(-daysSinceMonday*24*time.Hour - time.Duration(weeksAgo*7*24*uint64(time.Hour))).Unix())
}
