package calendar

import (
	"fmt"

	"github.com/ag7if/go-files"
	"github.com/fxtlabs/date"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar/kfm"
	"github.com/derhabicht/planning-tools/pkg/calendar/plancal"
)

func GenerateKFM(fy int) error {
	bd, err := date.ParseISO(config.GetString(config.Birthday))
	if err != nil {
		return errors.WithStack(err)
	}
	cal := plancal.NewCalendar(2026, bd)

	k := kfm.GenerateBlankKFM(cal)

	doc, err := yaml.Marshal(k)
	if err != nil {
		return errors.WithStack(err)
	}

	f, err := files.NewFile(fmt.Sprintf("PlanningCalendarKFM-FY%d.yaml", fy))
	if err != nil {
		return errors.WithStack(err)
	}

	err = f.WriteBytes(doc)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
