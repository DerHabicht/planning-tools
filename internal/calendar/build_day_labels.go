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

func BuildLabels(year, week int, outputFile files.File) error {
	bd, err := date.ParseISO(config.GetString(config.Birthday))
	if err != nil {
		return errors.WithStack(err)
	}
	cal := plancal.NewCalendar(year, bd)

	compiler, err := configureLaTeXCompiler()
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
