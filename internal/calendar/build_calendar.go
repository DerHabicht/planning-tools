package calendar

import (
	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/plancal"
	"github.com/derhabicht/planning-tools/reports/planning_calendar"
)

func generateCalendarLaTeX(cal calendar.Calendar, compiler *latex.Compiler, outputFile files.File) error {
	planningCal := planning_calendar.NewCalendar(cal)

	assets := []string{config.GetString(config.CoverLogo)}

	err := compiler.GenerateLaTeX(planningCal, outputFile, assets)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildCalendar(year int, outputFile files.File, logger logging.Logger) error {
	cal := plancal.NewCalendar(year)

	compiler, err := configureLaTeXCompiler(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateCalendarLaTeX(cal, compiler, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return err
	}

	return nil
}
