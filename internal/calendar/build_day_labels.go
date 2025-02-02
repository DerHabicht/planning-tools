package calendar

import (
	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/logging"
	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/plancal"
	"github.com/derhabicht/planning-tools/reports/planning_calendar"
)

func generateLabelLaTeX(cal calendar.Calendar, year, week int, compiler *latex.Compiler, outputFile files.File) error {
	labels, err := planning_calendar.NewDayLabels(cal, year, week)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.GenerateLaTeX(labels, outputFile, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildLabels(year, week int, outputFile files.File, logger logging.Logger) error {
	cal := plancal.NewCalendar(year)

	compiler, err := configureLaTeXCompiler(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateLabelLaTeX(cal, year, week, compiler, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return err
	}

	return nil
}
