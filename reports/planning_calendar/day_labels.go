package planning_calendar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"
)

type DayLabels struct {
	week calendar.Week
}

func NewDayLabels(cal calendar.Calendar, year, week int) (*DayLabels, error) {
	wk, err := cal.FetchWeek(year, week)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &DayLabels{week: wk}, nil
}

func (dl *DayLabels) generateWeekLabel(latex string) string {
	templ := templates.WeekLabelTemplate

	fy, fw := dl.week.FyWeek()
	cy, cw, cc := dl.week.ISOWeek()
	cyWeek := fmt.Sprintf("%sW%02d", cc.LaTeX(), cw)
	ft := dl.week.Trimester().Short()
	fq := dl.week.FiscalQuarter().Short()
	aq := dl.week.CalendarQuarter().Short()
	s := dl.week.Sprint().Short()

	templ = strings.ReplaceAll(templ, templates.CalendarYear1, strconv.Itoa(cy))
	templ = strings.ReplaceAll(templ, templates.FiscalTrimester, ft)
	templ = strings.ReplaceAll(templ, templates.FiscalQuarter, fq)
	templ = strings.ReplaceAll(templ, templates.FiscalWeek, fmt.Sprintf("W%02d", fw))
	templ = strings.ReplaceAll(templ, templates.CalendarQuarter, aq)
	templ = strings.ReplaceAll(templ, templates.Sprint, s)
	templ = strings.ReplaceAll(templ, templates.ISOWeek, cyWeek)
	templ = strings.ReplaceAll(templ, templates.Year, strconv.Itoa(fy))

	latex = strings.ReplaceAll(latex, templates.WeekData, templ)

	return latex
}

func (dl *DayLabels) generateDayLabels(latex string, day calendar.Day, idx int) string {
	const fullDateFormat = `Monday, 2 January 2006`
	const timeFormat = `1504`
	templ := templates.DayLabelTemplate

	fd := day.Date().Format(fullDateFormat)
	fy, fw := dl.week.FyWeek()
	cy, cw, cc := dl.week.ISOWeek()
	cyWeek := fmt.Sprintf("%sW%02d", cc.LaTeX(), cw)
	ft := dl.week.Trimester().Short()
	fq := dl.week.FiscalQuarter().Short()
	aq := dl.week.CalendarQuarter().Short()
	s := dl.week.Sprint().Short()
	ord := day.OrdinalDay()
	mjd := day.MJD()
	sr := day.Sunrise().Format(timeFormat)
	ss := day.Sunset().Format(timeFormat)

	templ = strings.ReplaceAll(templ, templates.FullDate, fd)
	templ = strings.ReplaceAll(templ, templates.CalendarYear1, strconv.Itoa(cy))
	templ = strings.ReplaceAll(templ, templates.FiscalTrimester, ft)
	templ = strings.ReplaceAll(templ, templates.FiscalQuarter, fq)
	templ = strings.ReplaceAll(templ, templates.FiscalWeek, fmt.Sprintf("W%02d", fw))
	templ = strings.ReplaceAll(templ, templates.CalendarQuarter, aq)
	templ = strings.ReplaceAll(templ, templates.Sprint, s)
	templ = strings.ReplaceAll(templ, templates.ISOWeek, cyWeek)
	templ = strings.ReplaceAll(templ, templates.OrdinalDay, fmt.Sprintf("%03d", ord))
	templ = strings.ReplaceAll(templ, templates.MJD, strconv.Itoa(mjd))
	templ = strings.ReplaceAll(templ, templates.SunriseTime, sr)
	templ = strings.ReplaceAll(templ, templates.SunsetTime, ss)
	templ = strings.ReplaceAll(templ, templates.Year, strconv.Itoa(fy))

	latex = strings.ReplaceAll(latex, templates.MonthDayData(idx), templ)

	return latex
}

func (dl *DayLabels) LaTeX() string {
	latex := templates.WeekLabels

	latex = dl.generateWeekLabel(latex)

	day := dl.week.StartDay()
	for i := 1; i <= 7; i++ {
		latex = dl.generateDayLabels(latex, day, i)
		day = day.Next()
	}

	return latex
}
