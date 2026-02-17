package planning_calendar

import (
	"strings"

	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"

	"github.com/derhabicht/planning-tools/pkg/calendar"
)

const quarterMonthCount = 3
const quarterWeekCount = 13

type Quarter struct {
	calendarQuarter calendar.Quarter
	fiscalQuarter   calendar.Quarter
	minimonths      []Minimonth
}

func NewQuarter(calendar, fiscal calendar.Quarter, minimonths map[string]Minimonth) Quarter {
	var mm []Minimonth

	mo := fiscal.FirstMonth()
	for i := 0; i < quarterMonthCount; i++ {
		mm = append(mm, minimonths[mo.Full()])
		mo = mo.Next()
	}

	return Quarter{
		calendarQuarter: calendar,
		fiscalQuarter:   fiscal,
		minimonths:      mm,
	}
}

func (q *Quarter) LaTeX() string {
	latex := templates.QuarterTemplate

	latex = strings.Replace(latex, templates.FullFiscalQuarter, q.fiscalQuarter.Full(), 1)
	latex = strings.Replace(latex, templates.FullCalendarQuarter, q.calendarQuarter.Full(), 1)

	for i, mm := range q.minimonths {
		latex = strings.Replace(latex, templates.MinimonthMacro(i+1), mm.LatexCommand(), 1)
	}

	return latex
}
