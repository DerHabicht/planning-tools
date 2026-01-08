package calendar

import (
	"fmt"

	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/ag7if"
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

func generateCardLaTeX(cal calendar.Calendar, year, week int, contexts []string, compiler *latex.Compiler, outputFile files.File) error {
	cards, err := planning_calendar.NewAG7IF5303(cal, year, week, contexts)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.GenerateLaTeX(cards, outputFile, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func buildLabels(year, week int, outputFile files.File) error {
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
		return errors.WithStack(err)
	}

	return nil
}

func buildCards(year, week int, contexts []string, outputFile files.File) error {
	bd, err := date.ParseISO(config.GetString(config.Birthday))
	if err != nil {
		return errors.WithStack(err)
	}
	cal := plancal.NewCalendar(year, bd)

	compiler, err := configureLaTeXCompiler()
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateCardLaTeX(cal, year, week, contexts, compiler, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildDL(year, period int, sprint, cards, labels bool, contexts []string) error {
	var weeks []int

	if sprint {
		spStr := fmt.Sprintf("S%02d", period)
		sp, _ := ag7if.ParseSprint(spStr)
		weeks = ag7if.ComputeSprintWeekNumbers(sp)
	} else {
		weeks = []int{period}
	}

	if cards {
		for _, w := range weeks {
			outputFile, err := files.NewFile(fmt.Sprintf("DayCards-%04dW%02d.pdf", year, w))
			if err != nil {
				return errors.WithMessage(err, "failed to create output file")
			}

			err = buildCards(year, w, contexts, outputFile)
			if err != nil {
				return errors.WithMessage(err, "failed to build cards")
			}
		}
	}

	if labels {
		for _, w := range weeks {
			outputFile, err := files.NewFile(fmt.Sprintf("DayLabels-%04dW%02d.pdf", year, w))
			if err != nil {
				return errors.WithMessage(err, "failed to create output file")
			}

			err = buildLabels(year, w, outputFile)
			if err != nil {
				return errors.WithMessage(err, "failed to build labels")
			}
		}
	}

	return nil
}
