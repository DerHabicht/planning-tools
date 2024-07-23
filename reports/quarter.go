package reports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/reports/templates"
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

func (q *Quarter) generateOKRHeader(latex string) string {
	hdr := templates.OKRHeaderTemplate

	w := q.calendarQuarter.FirstMonth().FirstWeek()
	for i := 1; i <= quarterWeekCount; i++ {
		_, wk, _ := w.ISOWeek()
		hdr = strings.Replace(hdr, fmt.Sprintf(templates.OKRHeaderWeekNumber, i), strconv.Itoa(wk), 1)
		w = w.Next()
	}

	latex = strings.Replace(latex, templates.OKRHeader, hdr, -1)

	return latex
}

func (q *Quarter) LaTeX() string {
	latex := templates.QuarterTemplate

	latex = strings.Replace(latex, templates.FullFiscalQuarter, q.fiscalQuarter.Full(), 1)
	latex = strings.Replace(latex, templates.FullCalendarQuarter, q.calendarQuarter.Full(), 1)

	for i, mm := range q.minimonths {
		latex = strings.Replace(latex, fmt.Sprintf(templates.MinimonthMacro, i+1), mm.LatexCommand(), 1)
	}

	latex = q.generateOKRHeader(latex)

	return latex
}
