package latex

import (
	"fmt"
	"os"
	"path/filepath"
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
	for month := 1; month <= 13; month++ {
		mt := NewMonthTemplate(ct.calendar, ct.monthTemplate)

		ct.calendarTemplate = strings.Replace(ct.calendarTemplate, fmt.Sprintf("+M%02d", month), mt.LaTeX(), 1)

		ct.calendar.NextMonth()
	}

	return ct.calendarTemplate
}
