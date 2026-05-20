package graph

import (
	"errors"
	"fmt"
	"time"
)

const (
	daysInWeek  = 7
	weeksInYear = 53
)

func GraphDateByCoordinates(x, y int) (time.Time, error) {
	if err := validateCoordinates(x, y); err != nil {
		return time.Time{}, err
	}

	startDate := getGraphStartDate()
	daysToAdd := x*daysInWeek + y

	return startDate.AddDate(0, 0, daysToAdd), nil
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

	// Откат до воскресенья (в начало графика)
	return t.AddDate(0, 0, -int(t.Weekday()))
}

func validateCoordinates(x, y int) error {
	var errs []error

	if y < 0 || y >= daysInWeek {
		errs = append(errs, fmt.Errorf("error days representation: %d", y))
	}

	if x < 0 || x >= weeksInYear {
		errs = append(errs, fmt.Errorf("error weeks representation: %d", x))
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errors.Join(errs...)
	}
}
