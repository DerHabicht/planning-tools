package planwx

import (
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/internal/clients"
	"github.com/derhabicht/planning-tools/pkg/metoc"
)

func FetchSunData(location metoc.Location, date metoc.Dtg, tzoffset int) (metoc.SunData, error) {
	resp, err := clients.FetchSunriseSunsetData(location, date)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	sunrise, err := metoc.ParseIsoDateTimeToDtg(resp.Results.Sunrise, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	sunset, err := metoc.ParseIsoDateTimeToDtg(resp.Results.Sunset, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	beginAt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.AstronomicalTwilightBegin, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	beginNt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.NauticalTwilightBegin, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	beginCt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.CivilTwilightBegin, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	endAt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.AstronomicalTwilightEnd, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	endNt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.NauticalTwilightEnd, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	endCt, err := metoc.ParseIsoDateTimeToDtg(resp.Results.CivilTwilightEnd, tzoffset)
	if err != nil {
		return metoc.SunData{}, errors.WithStack(err)
	}

	return metoc.SunData{
		Sunrise: sunrise,
		Sunset:  sunset,
		AstronomicalTwilight: metoc.Twilight{
			Begin: beginAt,
			End:   endAt,
		},
		NauticalTwilight: metoc.Twilight{
			Begin: beginNt,
			End:   endNt,
		},
		CivilTwilight: metoc.Twilight{
			Begin: beginCt,
			End:   endCt,
		},
	}, nil
}

func FetchMoonData(location metoc.Location, date metoc.Dtg, tzoffset int) (metoc.MoonData, error) {
	resp, err := clients.FetchUsnoData(location, date, tzoffset)
	if err != nil {
		return metoc.MoonData{}, errors.WithStack(err)
	}

	var rise metoc.Dtg
	var set metoc.Dtg

	moondata := resp.Properties.Data.MoonData
	for _, d := range moondata {
		if d.Phenomenon == "Rise" {
			rise, err = metoc.ParseTimeToDtg(d.Time, tzoffset)
			if err != nil {
				return metoc.MoonData{}, errors.WithStack(err)
			}
		} else if d.Phenomenon == "Set" {
			set, err = metoc.ParseTimeToDtg(d.Time, tzoffset)
			if err != nil {
				return metoc.MoonData{}, errors.WithStack(err)
			}
		}
	}

	var phase metoc.LunarPhase
	switch resp.Properties.Data.CurrentPhase {
	case "New Moon":
		phase = metoc.New
	case "Waxing Crescent":
		phase = metoc.WaxingCrescent
	case "First Quarter":
		phase = metoc.FirstQuarter
	case "Waxing Gibbous":
		phase = metoc.WaxingGibbous
	case "Full Moon":
		phase = metoc.Full
	case "Waning Gibbous":
		phase = metoc.WaningGibbous
	case "Third Quarter":
		phase = metoc.LastQuarter
	case "Waning Crescent":
		phase = metoc.WaningCrescent
	default:
		return metoc.MoonData{}, errors.Errorf(
			"failed to parse %s as a moon phase",
			resp.Properties.Data.CurrentPhase,
		)
	}

	return metoc.MoonData{
		Phase:    phase,
		MoonRise: rise,
		MoonSet:  set,
	}, nil
}

func FetchDailyAstroData(location metoc.Location, dtg metoc.Dtg, tzoffset int) (metoc.AstroData, error) {
	sunData, err := FetchSunData(location, dtg, tzoffset)
	if err != nil {
		return metoc.AstroData{}, errors.WithStack(err)
	}

	moonData, err := FetchMoonData(location, dtg, tzoffset)
	if err != nil {
		return metoc.AstroData{}, errors.WithStack(err)
	}

	return metoc.AstroData{SunData: sunData, MoonData: moonData}, nil
}

func FetchAstroDataForReport(report *metoc.MetocReport) error {
	for _, date := range report.Dates {
		ad, err := FetchDailyAstroData(report.Location, date, report.TzOffset)
		if err != nil {
			return errors.WithStack(err)
		}

		report.AstroData[date] = ad
	}

	return nil
}
