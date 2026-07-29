package render

import (
	"fmt"

	bimapLib "github.com/arkadashiim/go_gitify/internal/bitmap"
	"github.com/arkadashiim/go_gitify/internal/constants"
	"github.com/arkadashiim/go_gitify/internal/graph"
)

type PrintPoint = string

const (
	Empty   PrintPoint = "\x1b[38;5;236m▗▖\x1b[0m"
	Spotted PrintPoint = "\x1b[38;5;34m▗▖\x1b[0m"
)

var weekdayLabels = [constants.DaysInWeek]string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}

var monthAbbreviations = [...]string{
	"", // time.Month отсчитывается с 1 (Январь)
	"Янв", "Фев", "Мар", "Апр", "Май", "Июн",
	"Июл", "Авг", "Сен", "Окт", "Ноя", "Дек",
}

func RenderBitmap(bitmap bimapLib.Bitmap) {
	formatStringForSpace := "%-3s"

	fmt.Printf(formatStringForSpace, "")
	fmt.Println(monthHeader())

	for dayIndex, row := range bitmap {
		fmt.Printf(formatStringForSpace, weekdayLabels[dayIndex])

		for _, point := range row {
			fmt.Print(toPrintPoint(point))
		}

		fmt.Println()
	}
}

func toPrintPoint(p bimapLib.BitmapPoint) PrintPoint {
	if p == bimapLib.Spotted {
		return Spotted
	} else {
		return Empty
	}
}

type monthMark struct {
	week  int
	label []rune
}

func monthHeader() string {
	header := make([]rune, constants.WeeksInYear*2)
	for i := range header {
		header[i] = ' '
	}

	var marks []monthMark
	lastMonth := 0
	for week := range constants.WeeksInYear {
		date, err := graph.GraphDateByCoordinates(week, 0)
		if err != nil {
			continue
		}

		month := int(date.Month())
		if month == lastMonth {
			continue
		}
		lastMonth = month

		marks = append(marks, monthMark{week: week, label: []rune(monthAbbreviations[month])})
	}

	for markIndex, mark := range marks {
		available := len(header) - mark.week*2
		if markIndex+1 < len(marks) {
			available = (marks[markIndex+1].week - mark.week) * 2
		}

		if len(mark.label) > available {
			// соседний месяц наступает раньше, чем влезет подпись, поэтому скип
			continue
		}

		for i, r := range mark.label {
			header[mark.week*2+i] = r
		}
	}

	return string(header)
}
