package latex

import (
	"fmt"
	"strings"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

const dayDataTemplate = `+DY\\\moon{+FD}\\\vspace{2em}+YD\hfill`
const dayDataTemplateWithSolar = `+DY\\\moon{+FD}\\\vspace{1em}+SR\\+YD\hfill{}+SS`

type MonthTemplate struct {
	calendar calendar.Calendar
	month    string
	template string
}

func NewMonthTemplate(cal calendar.Calendar, template string) MonthTemplate {
	return MonthTemplate{
		calendar: cal,
		month:    cal.CurrentMonth(),
		template: template,
	}
}

func (m *MonthTemplate) fillDayData(dayStr string, dt date.Date) string {
	data := strings.Replace(dayDataTemplate, "+DY", dayStr, 1)
	data = strings.Replace(data, "+FD", dt.FormatISO(4), 1)
	data = strings.Replace(data, "+YD", fmt.Sprintf("%03d", dt.YearDay()), 1)

	return data
}

func (m *MonthTemplate) fillDayDataWithSolar(dayStr string, sunrise, sunset time.Time, dt date.Date) string {
	data := strings.Replace(dayDataTemplateWithSolar, "+DY", dayStr, 1)
	data = strings.Replace(data, "+SR", sunrise.Format("1504"), 1)
	data = strings.Replace(data, "+SS", sunset.Format("1504"), 1)
	data = strings.Replace(data, "+FD", dt.FormatISO(4), 1)
	data = strings.Replace(data, "+YD", fmt.Sprintf("%03d", dt.YearDay()), 1)

	return data
}

func (m *MonthTemplate) fillWeekData(weekNum int, week calendar.Week) {
	m.template = strings.Replace(m.template, fmt.Sprintf("+FT%d", weekNum), week.FyTrimester(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+FQ%d", weekNum), week.FyQuarter(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+FW%d", weekNum), week.FyWeek(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+AQ%d", weekNum), week.Ag7ifQuarter(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+AS%d", weekNum), week.Ag7ifSprint(), 1)
	m.template = strings.Replace(m.template, fmt.Sprintf("+IW%d", weekNum), week.IsoWeek(), 1)
}

func (m *MonthTemplate) LaTeX() string {
	m.template = strings.Replace(m.template, "+M", m.calendar.CurrentMonth(), 1)

	day := 0
	week := m.calendar.CurrentWeek()

	for wk := 1; wk <= 6; wk++ {
		week.Reset()
		m.fillWeekData(wk, week)

		for i := 0; i < 7; i++ {
			var dayStr string
			var err error
			if day == 0 {
				dayStr, err = week.CurrentDayStr(true)
			} else {
				dayStr, err = week.CurrentDayStr(false)
			}
			if err != nil {
				panic(fmt.Errorf("unexpectedly got to the end of week %s", week.IsoWeek()))
			}

			currentDate, _ := week.CurrentDay()

			sunrise, sunset, err := week.CurrentDaySunriseSunset()
			var dayData string
			if err != nil {
				dayData = m.fillDayData(dayStr, currentDate)
			} else {
				dayData = m.fillDayDataWithSolar(dayStr, sunrise, sunset, currentDate)
			}

			day += 1
			m.template = strings.Replace(m.template, fmt.Sprintf("+D%02d", day), dayData, 1)

			week.Next()
			/*
				if err != nil {
					panic(fmt.Errorf("unexpectedly got to the end of week %s", week.IsoWeek()))
				}
			*/
		}

		week = m.calendar.NextWeek()
	}

	return m.template
}
