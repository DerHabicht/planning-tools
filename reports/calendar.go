package reports

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/internal/config"
)

type Calendar struct {
	calendar          calendar.Calendar
	solsticeTable     SolsticeTable
	calendarTemplate  string
	trimesterTemplate string
	quarterTemplate   string
	monthTemplate     string
}

func NewCalendar(cal calendar.Calendar) (Calendar, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}

	raw, err := os.ReadFile(filepath.Join(cfgDir, "assets", "calendar.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	calendarTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "trimester.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	trimesterTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "quarter.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	quarterTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "month.tex"))
	if err != nil {
		return Calendar{}, errors.WithStack(err)
	}
	monthTemplate := string(raw)

	return Calendar{
		calendar:          cal,
		solsticeTable:     NewSolsticeTable(cal.SolsticeTable()),
		calendarTemplate:  calendarTemplate,
		trimesterTemplate: trimesterTemplate,
		quarterTemplate:   quarterTemplate,
		monthTemplate:     monthTemplate,
	}, nil
}

func (ct Calendar) LaTeX() string {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		panic(err)
	}
	coverLogoPath := filepath.Join(cfgDir, "assets", config.GetString("cover_logo"))

	holidayList := NewHolidayList(ct.calendar)

	latex := ct.calendarTemplate

	latex = strings.Replace(latex, "+PIC", coverLogoPath, 1)
	latex = strings.Replace(latex, "+CAL_START", fmt.Sprintf("October %d", ct.calendar.FiscalYear()-1), 1)
	latex = strings.Replace(latex, "+CAL_END", fmt.Sprintf("December %d", ct.calendar.FiscalYear()), 1)
	latex = strings.Replace(latex, "+JP_START", fmt.Sprintf("%d", ct.calendar.StartingJulianPeriod()), 1)
	latex = strings.Replace(latex, "+JP_END", fmt.Sprintf("%d", ct.calendar.StartingJulianPeriod()+1), 1)
	latex = strings.Replace(latex, "+LCD", config.GetString("lunar_calibration_date"), 1)
	latex = strings.Replace(latex, "+CY1", fmt.Sprintf("%d", ct.calendar.FiscalYear()-1), 2)
	latex = strings.Replace(latex, "+CY2", fmt.Sprintf("%d", ct.calendar.FiscalYear()), 2)
	latex = strings.Replace(latex, "+ABBVS", holidayList.LaTeX(), 1)
	latex = strings.Replace(latex, "+SOLSTICES", ct.solsticeTable.LaTeX(), 1)

	fy := ct.calendar.FiscalYear() - 1
	tri := calendar.FyT3
	for trimester := 0; trimester <= 4; trimester++ {
		tt := NewTrimesterTemplate(tri, fy, ct.trimesterTemplate)

		latex = strings.Replace(latex, fmt.Sprintf("+T%d", trimester), tt.LaTeX(), 1)

		fy, tri = tri.NextTrimester(fy)
	}

	fy = ct.calendar.FiscalYear() - 1
	fyQtr := calendar.FyQ4
	for quarter := 0; quarter <= 5; quarter++ {
		qt := NewQuarterTemplate(fyQtr, fy, ct.quarterTemplate)

		latex = strings.Replace(latex, fmt.Sprintf("+Q%d", quarter), qt.LaTeX(), 1)

		fy, fyQtr = fyQtr.NextQuarter(fy)
	}

	for month := 1; month <= 15; month++ {
		mt := NewMonthTemplate(ct.calendar, ct.monthTemplate)

		latex = strings.Replace(latex, fmt.Sprintf("+M%02d", month), mt.LaTeX(), 1)

		ct.calendar.NextMonth()
	}

	return latex
}
