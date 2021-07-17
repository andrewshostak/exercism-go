package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(week WeekSchedule, weekDay time.Weekday, month time.Month, year int) int {
	days := []int{}
	lastDayOfMonth := getLastDayOfMonth(month, year)
	for day := 1; day <= lastDayOfMonth; day++ {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == weekDay {
			days = append(days, date.Day())
		}
	}

	switch week {
	case First:
		return days[0]
	case Second:
		return days[1]
	case Third:
		return days[2]
	case Fourth:
		return days[3]
	case Last:
		return days[len(days)-1:][0]
	case Teenth:
		for _, v := range days {
			if v > 12 && v < 20 {
				return v
			}
		}
	}

	return 0
}

func getLastDayOfMonth(month time.Month, year int) int {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDay.Day()
}
