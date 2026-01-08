package planning_calendar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"
)

type AG7IF5303 struct {
	week     calendar.Week
	contexts []string
}

func NewAG7IF5303(cal calendar.Calendar, year, week int, contexts []string) (*AG7IF5303, error) {
	wk, err := cal.FetchWeek(year, week)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AG7IF5303{week: wk, contexts: contexts}, nil
}

func (f *AG7IF5303) generateContext(latex, context string, day calendar.Day) string {
	const fullDateFormat = `Monday, 2 January 2006`
	const dateFormat = "2006-01-02"
	const timeFormat = "1504"
	templ := templates.DayCardData

	fd := day.Date().Format(fullDateFormat)
	isodate := day.Date().Format(dateFormat)
	fy, fw := f.week.FyWeek()
	cy, cw, cc := f.week.ISOWeek()
	cyWeek := fmt.Sprintf("%sW%02d", cc.LaTeX(), cw)
	ft := f.week.Trimester().Short()
	fq := f.week.FiscalQuarter().Short()
	aq := f.week.CalendarQuarter().Short()
	s := f.week.Sprint().Short()
	ord := day.OrdinalDay()
	mjd := day.MJD()
	sr := day.Sunrise().Format(timeFormat)
	ss := day.Sunset().Format(timeFormat)

	latex = strings.ReplaceAll(latex, templates.Context, context)
	latex = strings.ReplaceAll(latex, templates.ISODate, isodate)

	templ = strings.ReplaceAll(templ, templates.FullDate, fd)

	templ = strings.ReplaceAll(templ, templates.FiscalTrimester, ft)
	templ = strings.ReplaceAll(templ, templates.FiscalQuarter, fq)
	templ = strings.ReplaceAll(templ, templates.OrdinalDay, fmt.Sprintf("%03d", ord))

	templ = strings.ReplaceAll(templ, templates.FiscalWeek, fmt.Sprintf("W%02d", fw))
	templ = strings.ReplaceAll(templ, templates.MJD, strconv.Itoa(mjd))

	templ = strings.ReplaceAll(templ, templates.CalendarYear1, strconv.Itoa(cy))
	templ = strings.ReplaceAll(templ, templates.CalendarQuarter, aq)
	templ = strings.ReplaceAll(templ, templates.Sprint, s)
	templ = strings.ReplaceAll(templ, templates.SunriseTime, sr)

	templ = strings.ReplaceAll(templ, templates.ISOWeek, cyWeek)
	templ = strings.ReplaceAll(templ, templates.SunsetTime, ss)

	latex = strings.ReplaceAll(latex, templates.DayLabel, templ)

	// Has to be executed last to avoid premature matches (e.g. w/+YD)
	templ = strings.ReplaceAll(templ, templates.Year, strconv.Itoa(fy))

	return latex
}

func (f *AG7IF5303) generateDayCard(latex string, day calendar.Day, idx int) string {
	cards := ""
	var contexts []string
	if f.contexts == nil {
		contexts = []string{"ALL"}
	} else {
		contexts = f.contexts
	}

	for i, ctx := range contexts {
		if i > 0 {
			cards += `\pagebreak
\setcounter{page}{1}
`
		}
		cards += f.generateContext(templates.AG7IF5303, ctx, day)
	}

	latex = strings.ReplaceAll(latex, templates.DayCard(idx), cards)

	return latex
}

func (f *AG7IF5303) LaTeX() string {
	latex := templates.DayCards

	day := f.week.StartDay()
	for i := 1; i <= 7; i++ {
		latex = f.generateDayCard(latex, day, i)
		day = day.Next()
	}

	return latex
}
