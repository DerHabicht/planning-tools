package reports

import (
	"fmt"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

type TrimesterTemplate struct {
	fiscalYear         int
	trimester          calendar.FyTrimester
	startDate          date.Date
	miniMonthTemplates []miniMonthTemplate
	template           string
}

func NewTrimesterTemplate(trimester calendar.FyTrimester, fy int, template string, miniMonthTemplates []miniMonthTemplate) TrimesterTemplate {
	return TrimesterTemplate{
		fiscalYear:         fy,
		trimester:          trimester,
		startDate:          trimester.StartDate(fy),
		miniMonthTemplates: miniMonthTemplates,
		template:           template,
	}
}

func (t TrimesterTemplate) LaTeX() string {
	template := strings.Replace(t.template, "+T", t.trimester.FullName(t.fiscalYear), 1)

	for i, v := range t.miniMonthTemplates {
		key := fmt.Sprintf("+M%dCMD", i+1)
		template = strings.Replace(template, key, v.LaTeXCommand(), 1)
	}

	return template
}
