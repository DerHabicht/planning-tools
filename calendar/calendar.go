package calendar

import (
	"fmt"
	"strings"
	"time"

	"github.com/fxtlabs/date"
	"github.com/nathan-osman/go-sunrise"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/internal/config"
)

func computeNearestMonday(d date.Date) date.Date {
	var dd int
	wd := d.Weekday()
	switch {
	case time.Tuesday <= wd && wd <= time.Thursday:
		dd = -1 * (int(wd) - 1)
	case wd >= time.Friday:
		dd = 8 - int(wd)
	case wd == time.Sunday:
		dd = 1
	default:
		dd = 0
	}

	return d.Add(dd)
}

func computerNearestThursday(d date.Date) date.Date {
	monday := computeNearestMonday(d)
	return monday.Add(3)
}

func ComputeLastDayOfMonth(d date.Date) int {
	switch d.Month() {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if date.New(d.Year(), time.December, 31).YearDay() == 366 {
			return 29
		}

		return 28
	default:
		panic(errors.Errorf("not a valid time.Month value: %d", d.Month()))
	}
}

type EndOfWeekError struct{}

func (eow EndOfWeekError) Error() string {
	return "reached the end of the week"
}

type Week struct {
	monday          date.Date
	fiscalYear      CapFiscalYear
	fyTrimester     FyTrimester
	fyQuarter       FyQuarter
	ag7ifQuarter    Ag7ifQuarter
	ag7ifSprint     Ag7ifSprint
	isoWeek         int
	card            Card
	currentDay      date.Date
	location        *time.Location
	holidayCalendar *HolidayCalendar
	solsticeTable   SolsticeTable
}

func NewWeek(d date.Date, loc *time.Location, hc *HolidayCalendar, st SolsticeTable) *Week {
	_, isoWeek := d.ISOWeek()

	fiscalYear := FiscalYearFromDate(d)

	var monday date.Date
	if d.Weekday() == time.Sunday {
		monday = d.Add(-6)
	} else {
		monday = d.Add(-1 * (int(d.Weekday()) - 1))
	}

	thursday := monday.Add(3)

	fyTrimester := computeCapTrimester(thursday)
	fyQuarter := computeCapQuarter(thursday)
	ag7ifQuarter := computeAg7ifQuarter(isoWeek)
	ag7ifSprint := computeAg7ifSprint(isoWeek)

	return &Week{
		monday:          monday,
		fiscalYear:      fiscalYear,
		fyTrimester:     fyTrimester,
		fyQuarter:       fyQuarter,
		ag7ifQuarter:    ag7ifQuarter,
		ag7ifSprint:     ag7ifSprint,
		isoWeek:         isoWeek,
		card:            GetWeekCard(isoWeek),
		location:        loc,
		holidayCalendar: hc,
		solsticeTable:   st,
	}
}

func (w *Week) CurrentDay() (date.Date, error) {
	if w.currentDay.Sub(w.monday) >= 7 {
		return w.monday, EndOfWeekError{}
	}

	return w.currentDay, nil
}

func (w *Week) CurrentDayStr(month bool) (string, error) {
	if w.currentDay.Sub(w.monday) >= 7 {
		return "", EndOfWeekError{}
	}

	if month || w.currentDay.Day() == 1 {
		return fmt.Sprintf("%s %d",
			strings.ToUpper(w.currentDay.Month().String())[:3],
			w.currentDay.Day(),
		), nil
	}

	return fmt.Sprintf("%d", w.currentDay.Day()), nil
}

func (w *Week) CurrentDaySunriseSunset() (time.Time, time.Time, error) {
	homeLat := config.GetFloat64("home_location.lat")
	homeLong := config.GetFloat64("home_location.long")

	sr, ss := sunrise.SunriseSunset(homeLat, homeLong, w.currentDay.Year(), w.currentDay.Month(), w.currentDay.Day())

	return sr.In(w.location), ss.In(w.location), nil
}

func (w *Week) IsCurrentDaySolstice() (bool, Solstice) {
	return w.solsticeTable.IsSolstice(w.currentDay)
}

func (w *Week) Next() (date.Date, error) {
	w.currentDay = w.currentDay.Add(1)

	if w.currentDay.Sub(w.monday) >= 7 {
		return w.monday, EndOfWeekError{}
	}

	return w.currentDay, nil
}

func (w *Week) Reset() date.Date {
	w.currentDay = w.monday
	return w.currentDay
}

func (w *Week) FyTrimester() string {
	return w.fyTrimester.String()
}

func (w *Week) FyQuarter() string {
	return w.fyQuarter.String()
}

func (w *Week) FyWeek() string {
	wk, err := w.fiscalYear.FyWeek(w.monday)
	if err != nil {
		panic(fmt.Errorf("CAP Fiscal Year was improperly initialized for week %d: %s", w.isoWeek, err))
	}

	return fmt.Sprintf("W%02d", wk)
}

func (w *Week) Ag7ifQuarter() string {
	return w.ag7ifQuarter.String()
}

func (w *Week) Ag7ifSprint() string {
	return w.ag7ifSprint.String()
}

func (w *Week) IsoWeek() (string, Card) {
	return fmt.Sprintf("W%02d", w.isoWeek), w.card
}

func (w *Week) IsCurrentDayHoliday() (bool, bool, *Holiday) {
	return w.holidayCalendar.IsHoliday(w.currentDay.Local())
}

type Calendar struct {
	fiscalYear          int
	startJulianPeriod   int
	currentCalendarYear int
	currentDoomsday     time.Weekday
	currentMonth        time.Month
	currentWeek         *Week
	location            *time.Location
	holidayCalendar     *HolidayCalendar
	solsticeTable       SolsticeTable
}

func NewCalendar(fiscalYear int) Calendar {
	loc, err := time.LoadLocation(config.GetString("home_location.tz"))
	if err != nil {
		loc = time.UTC
	}

	hc := NewHolidayCalendar()

	c := Calendar{
		fiscalYear:          fiscalYear,
		startJulianPeriod:   (fiscalYear - 1) + 4713,
		currentCalendarYear: fiscalYear - 1,
		location:            loc,
		holidayCalendar:     &hc,
		solsticeTable:       NewSolsticeTable(fiscalYear),
	}

	c.initMonthAndWeek()

	return c
}

func (c *Calendar) initMonthAndWeek() {
	c.currentCalendarYear = c.fiscalYear - 1
	c.currentDoomsday = ComputeDoomsday(c.currentCalendarYear)
	c.currentMonth = time.October
	c.currentWeek = NewWeek(date.New(c.currentCalendarYear, time.October, 1), c.location, c.holidayCalendar, c.solsticeTable)
}

func (c *Calendar) CurrentWeek() *Week {
	return c.currentWeek
}

func (c *Calendar) FiscalYear() int {
	return c.fiscalYear
}

func (c *Calendar) SolsticeTable() SolsticeTable {
	return c.solsticeTable
}

func (c *Calendar) StartingJulianPeriod() int {
	return c.startJulianPeriod
}

func (c *Calendar) NextWeek() *Week {
	d := c.currentWeek.Reset()

	c.currentWeek = NewWeek(d.Add(7), c.location, c.holidayCalendar, c.solsticeTable)

	return c.currentWeek
}

func (c *Calendar) CurrentMonth() (int, time.Month) {
	return c.currentCalendarYear, c.currentMonth
}

func (c *Calendar) CurrentMonthStr() string {
	return fmt.Sprintf("%s %d", c.currentMonth, c.currentCalendarYear)
}

func (c *Calendar) NextMonth() (string, *Week) {
	switch {
	case c.currentMonth == time.December:
		c.currentMonth = time.January
		c.currentCalendarYear += 1
		c.currentDoomsday = ComputeDoomsday(c.currentCalendarYear)
	default:
		c.currentMonth += 1
	}

	c.currentWeek = NewWeek(date.New(c.currentCalendarYear, c.currentMonth, 1), c.location, c.holidayCalendar, c.solsticeTable)
	return c.CurrentMonthStr(), c.currentWeek
}

func (c *Calendar) Reset() (string, *Week) {
	c.initMonthAndWeek()

	return c.CurrentMonthStr(), c.currentWeek
}

func (c *Calendar) Holidays() []Holiday {
	holidays := make([]Holiday, 0)

	for _, holiday := range c.holidayCalendar.holidays {
		holidays = append(holidays, holiday)
	}

	return holidays
}
