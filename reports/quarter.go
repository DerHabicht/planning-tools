package reports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

const okrHeaderTemplate = `Objectives & Key Results & +W01 & +W02 & +W03 & +W04 & +W05 & +W06 & +W07 & +W08 & +W09 & +W10 & +W11 & +W12 & +W13`

type QuarterTemplate struct {
	fiscalYear         int
	calendarYear       int
	fiscalQuarter      calendar.FyQuarter
	calQuarter         calendar.Ag7ifQuarter
	startDate          date.Date
	miniMonthTemplates []miniMonthTemplate
	template           string
}

func NewQuarterTemplate(quarter calendar.FyQuarter, fy int, template string, miniMonthTemplates []miniMonthTemplate) QuarterTemplate {
	calendarYear, calQuarter := quarter.CyQuarter(fy)

	return QuarterTemplate{
		fiscalYear:         fy,
		calendarYear:       calendarYear,
		fiscalQuarter:      quarter,
		calQuarter:         calQuarter,
		startDate:          quarter.StartDate(fy),
		miniMonthTemplates: miniMonthTemplates,
		template:           template,
	}
}

func (q QuarterTemplate) setOKRHeader(template string) string {
	header := okrHeaderTemplate
	week := q.calQuarter.StartWeek()
	for w := 1; w <= 13; w++ {
		header = strings.Replace(header, fmt.Sprintf("+W%02d", w), strconv.Itoa(week), 1)
		week++
	}

	return strings.Replace(template, "+OKRHDR", header, -1)
}

func (q QuarterTemplate) LaTeX() string {
	template := strings.Replace(q.template, "+CYQ", q.calQuarter.FullName(q.calendarYear), 1)
	template = strings.Replace(template, "+FYQ", q.fiscalQuarter.FullName(q.fiscalYear), 1)
	template = q.setOKRHeader(template)

	for i, v := range q.miniMonthTemplates {
		key := fmt.Sprintf("+M%dCMD", i+1)
		template = strings.Replace(template, key, v.LaTeXCommand(), 1)
	}

	return template
}
