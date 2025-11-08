package calendar

import (
	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/plancal"
	"github.com/derhabicht/planning-tools/reports/planning_calendar"
)

func generateTabsLaTeX(cal calendar.Calendar, startTab int, compiler *latex.Compiler, outputFile files.File) error {
	tabs := planning_calendar.NewCalendarTabs(cal, startTab)

	err := compiler.GenerateLaTeX(tabs, outputFile, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildCalendarTabs(year, startTab int, outputFile files.File) error {
	bd, err := date.ParseISO(config.GetString(config.Birthday))
	if err != nil {
		return errors.WithStack(err)
	}
	cal := plancal.NewCalendar(year, bd)

	compiler, err := configureLaTeXCompiler()
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateTabsLaTeX(cal, startTab, compiler, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
