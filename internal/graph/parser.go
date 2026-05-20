package graph

import "time"

const daysInWeek = 7

func GraphDateByCoordinates(x, y int) time.Time {
	startDate := getGraphStartDate()
	daysToAdd := x*daysInWeek + y

	return startDate.AddDate(0, 0, daysToAdd)
}

func getGraphStartDate() time.Time {
	// Дата год назд
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	t := time.Date(
		oneYearAgo.Year(),
		oneYearAgo.Month(),
		oneYearAgo.Day(),
		0, 0, 0, 0,
		time.UTC,
	)

	// Откат до воскресенье (в начало графика)
	return t.AddDate(0, 0, -int(t.Weekday()))
}
