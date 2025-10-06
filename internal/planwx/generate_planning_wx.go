package planwx

import (
	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/internal/logging"
	"github.com/derhabicht/planning-tools/pkg/metoc"
	"github.com/derhabicht/planning-tools/reports/metoc_report"
)

func parsePlanFile(planFile files.File) (metoc.Plan, error) {
	raw, err := planFile.ReadFile()
	if err != nil {
		return metoc.Plan{}, errors.WithStack(err)
	}

	plan := metoc.Plan{}
	err = yaml.Unmarshal(raw, &plan)
	if err != nil {
		return metoc.Plan{}, errors.WithStack(err)
	}

	return plan, nil
}

func loadPlanLocationDates(planLocation metoc.PlanLocation, report *metoc.MetocReport) error {
	for _, d := range planLocation.Dates {
		dtg, err := metoc.ParseIsoDateToDtg(d)
		if err != nil {
			return errors.WithStack(err)
		}

		report.AddDateToMetocReport(dtg)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func fetchReports(plan metoc.Plan) ([]metoc.MetocReport, error) {
	var reports []metoc.MetocReport
	for _, l := range plan.Locations {
		r, err := metoc.NewMetocReport(l.Name, l.Mgrs, plan.Tzoffset)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = loadPlanLocationDates(l, &r)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		reports = append(reports, r)
	}

	return reports, nil
}

func buildReport(reports []metoc.MetocReport, tzOffset int) (*metoc_report.PlanWxReport, error) {
	planwx := metoc_report.NewPlanWxReport(metoc.DtgNow(tzOffset))

	for _, report := range reports {
		planwx.AddLocation(report.Location.Mgrs(), report.Location.Name, report.Location.Mgrs())

		err := FetchAstroDataForReport(&report)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = FetchWeatherDataForReport(&report)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, date := range report.Dates {
			planwx.AddAstroData(report.Location.Mgrs(), date, report.AstroData[date])
			apf := EncodeDailyForecast(report.Forecast[date], report.Location, report.Generated)
			planwx.AddWxData(report.Location.Mgrs(), date, report.Forecast[date], apf)
		}
	}

	return planwx, nil
}

func configureLaTeXCompiler() (*latex.Compiler, error) {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find cache directory")
	}

	compiler := latex.NewCompiler(latex.XeLaTeX, latex.NoBib, *cacheDir)

	return &compiler, nil
}

func Generate(planFile, outputFile files.File, logger logging.Logger) error {
	plan, err := parsePlanFile(planFile)
	if err != nil {
		return errors.WithStack(err)
	}

	reports, err := fetchReports(plan)
	if err != nil {
		return errors.WithStack(err)
	}

	planwx, err := buildReport(reports, plan.Tzoffset)
	if err != nil {
		return errors.WithStack(err)
	}

	compiler, err := configureLaTeXCompiler()
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.GenerateLaTeX(planwx, outputFile, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
