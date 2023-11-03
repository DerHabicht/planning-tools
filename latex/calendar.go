package latex

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/config"
)

type CalendarTemplate struct {
	calendar         calendar.Calendar
	calendarTemplate string
	monthTemplate    string
}

func NewCalendarTemplate(cal calendar.Calendar) (CalendarTemplate, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return CalendarTemplate{}, errors.WithStack(err)
	}

	raw, err := os.ReadFile(filepath.Join(cfgDir, "assets", "calendar.tex"))
	if err != nil {
		return CalendarTemplate{}, errors.WithStack(err)
	}
	calendarTemplate := string(raw)

	raw, err = os.ReadFile(filepath.Join(cfgDir, "assets", "month.tex"))
	if err != nil {
		return CalendarTemplate{}, errors.WithStack(err)
	}
	monthTemplate := string(raw)

	return CalendarTemplate{
		calendar:         cal,
		calendarTemplate: calendarTemplate,
		monthTemplate:    monthTemplate,
	}, nil
}

func (ct *CalendarTemplate) LaTeX() string {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		panic(err)
	}
	coverLogoPath := filepath.Join(cfgDir, "assets", config.GetString("cover_logo"))

	holidayList := NewHolidayList(ct.calendar)

	latex := ct.calendarTemplate

	latex = strings.Replace(latex, "+PIC", coverLogoPath, 1)
	latex = strings.Replace(latex, "+FY", strconv.Itoa(ct.calendar.FiscalYear()), 1)
	latex = strings.Replace(latex, "+LCD", config.GetString("lunar_calibration_date"), 1)
	latex = strings.Replace(latex, "+CY1", fmt.Sprintf("%d", ct.calendar.FiscalYear()-1), 2)
	latex = strings.Replace(latex, "+CY2", fmt.Sprintf("%d", ct.calendar.FiscalYear()), 2)
	latex = strings.Replace(latex, "+ABBVS", holidayList.LaTeX(), 1)

	for month := 1; month <= 13; month++ {
		mt := NewMonthTemplate(ct.calendar, ct.monthTemplate)

		latex = strings.Replace(latex, fmt.Sprintf("+M%02d", month), mt.LaTeX(), 1)

		ct.calendar.NextMonth()
	}

	return latex
}
