package planning_calendar

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/doomsday"
)

const monthWeekCount = 6
const dayCount = 7

type Month struct {
	calendar  calendar.Calendar
	month     calendar.Month
	prevMonth Minimonth
	nextMonth Minimonth
}

func NewMonth(calendar calendar.Calendar, month calendar.Month, minimonths map[string]Minimonth) *Month {
	prev := month.Prev()
	next := month.Next()

	return &Month{
		calendar:  calendar,
		month:     month,
		prevMonth: minimonths[prev.Full()],
		nextMonth: minimonths[next.Full()],
	}
}

func (m *Month) generateMinimonths(latex string) string {
	latex = strings.Replace(latex, templates.MinimonthPrevious, m.prevMonth.LatexCommand(), 1)
	latex = strings.Replace(latex, templates.MinimonthNext, m.nextMonth.LatexCommand(), 1)

	return latex
}

func (m *Month) generateWeekdayHeader(latex string) string {
	header := templates.MonthWeekdayHeaderTemplate

	year := m.month.Year()
	dd := doomsday.ComputeDoomsday(year)

	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		abbv := strings.ToUpper(weekday.String())[:3]

		repl := abbv
		if weekday == dd {
			repl = fmt.Sprintf("<%s>", abbv)
		}

		header = strings.Replace(header, fmt.Sprintf("+%s", abbv), repl, 1)
	}

	latex = strings.Replace(latex, templates.WeekdayHeader, header, 1)

	return latex
}

func (m *Month) generateWeekData(week calendar.Week, latex string) string {
	_, fyWeek := week.FyWeek()
	_, cyWeek, card := week.ISOWeek()

	latex = strings.Replace(latex, templates.FiscalTrimester, week.Trimester().Short(), 1)
	latex = strings.Replace(latex, templates.FiscalQuarter, week.FiscalQuarter().Short(), 1)
	latex = strings.Replace(latex, templates.FiscalWeek, fmt.Sprintf("W%02d", fyWeek), 1)
	latex = strings.Replace(latex, templates.CalendarQuarter, week.CalendarQuarter().Short(), 1)
	latex = strings.Replace(latex, templates.Sprint, week.Sprint().Short(), 1)

	cyWeekStr := `\colorbox{%s}{\textcolor{white}{%sW%02d}}`
	if cyWeek%2 == 0 {
		latex = strings.Replace(latex, templates.ISOWeek, fmt.Sprintf(cyWeekStr, "blue", card.LaTeX(), cyWeek), 1)
	} else {
		latex = strings.Replace(latex, templates.ISOWeek, fmt.Sprintf(cyWeekStr, "red", card.LaTeX(), cyWeek), 1)
	}

	week = week.Next()

	return latex
}

func (m *Month) generateDayData(week calendar.Week, latex string, firstWeek bool) string {
	const timeFormat = `1504`

	d := week.StartDay()
	for i := 1; i <= dayCount; i++ {
		day := templates.MonthDayTemplate

		dayStr := strconv.Itoa(d.Date().Day())
		if firstWeek || d.Date().Day() == 1 {
			if firstWeek {
				firstWeek = false
			}
			dayStr = strings.ToUpper(d.Date().Month().String())[:3] + " " + dayStr
		}

		solstice := d.IsSolstice()
		if solstice != calendar.NoSolstice {
			dayStr = fmt.Sprintf(`%s\hfill{}%s`, solstice.LaTeX(), dayStr)
		}

		day = strings.Replace(day, templates.MonthDay, dayStr, 1)

		act, obs, holiday := d.IsHoliday()
		if act {
			day = strings.Replace(day, templates.Holiday, fmt.Sprintf(`%s\hfill{}`, holiday), 1)
		} else if obs {
			day = strings.Replace(day, templates.Holiday, fmt.Sprintf(`%s*\hfill{}`, holiday), 1)
		} else {
			day = strings.Replace(day, templates.Holiday, "", 1)
		}

		// This logic marks the Sunday square according to the typical second hour schedule for Sunday meetings
		// of the Church of Jesus Christ of Latter-day Saints. 1st and 3rd Sundays are Sunday School (
		// SS for Sontagsschule) and the 2nd and 4th Sundays are for quorum or class (
		// KK for Kollegium/Klasse) meetings. 5th Sundays are always Bishopric discretion and, therefore,
		// just marked as "5S" for "5. Sontag" (5th Sunday).
		if d.Weekday() == time.Sunday {
			switch d.WeekdayOccurrenceInMonth() {
			case 1, 3:
				day = strings.Replace(day, templates.Sunday, `\hfill{}SS\\`, 1)
			case 2, 4:
				day = strings.Replace(day, templates.Sunday, `\hfill{}KK\\`, 1)
			default:
				day = strings.Replace(day, templates.Sunday, `\hfill{}5S\\`, 1)
			}
		} else {
			day = strings.Replace(day, templates.Sunday, `\vspace{1em}`, 1)
		}

		day = strings.Replace(day, templates.FullDate, d.ISODate(), 1)
		day = strings.Replace(day, templates.OrdinalDay, fmt.Sprintf("%03d", d.OrdinalDay()), 1)
		day = strings.Replace(day, templates.SunriseTime, d.Sunrise().Format(timeFormat), 1)
		day = strings.Replace(day, templates.MJD, strconv.Itoa(d.MJD()), 1)
		day = strings.Replace(day, templates.SunsetTime, d.Sunset().Format(timeFormat), 1)

		latex = strings.Replace(latex, templates.MonthDayData(i), day, 1)

		d = d.Next()
	}

	return latex
}

func (m *Month) generateWeeks(latex string) string {
	week := m.month.FirstWeek()
	for i := 1; i <= monthWeekCount; i++ {
		s := templates.MonthWeekTemplate
		s = m.generateWeekData(week, s)
		s = m.generateDayData(week, s, i == 1)
		latex = strings.Replace(latex, templates.MonthPageWeek(i), s, 1)
		week = week.Next()
	}

	return latex
}

func (m *Month) LaTeX() string {
	latex := templates.MonthTemplate

	latex = strings.Replace(latex, templates.MonthNameFull, m.month.Full(), 1)
	latex = m.generateMinimonths(latex)
	latex = m.generateWeekdayHeader(latex)
	latex = m.generateWeeks(latex)

	return latex
}
