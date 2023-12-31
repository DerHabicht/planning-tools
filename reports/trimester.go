package reports

import (
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

type TrimesterTemplate struct {
	fiscalYear int
	trimester  calendar.FyTrimester
	startDate  date.Date
	template   string
}

func NewTrimesterTemplate(trimester calendar.FyTrimester, fy int, template string) TrimesterTemplate {
	return TrimesterTemplate{
		fiscalYear: fy,
		trimester:  trimester,
		startDate:  trimester.StartDate(fy),
		template:   template,
	}
}

func (t TrimesterTemplate) LaTeX() string {
	template := strings.Replace(t.template, "+T", t.trimester.FullName(t.fiscalYear), 1)

	template = fillMiniCalMonthNames(template, t.startDate, 4)
	template = fillMiniCalMonths(template, t.startDate, 4)

	return template
}
