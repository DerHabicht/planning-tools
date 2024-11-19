package planwx

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/derhabicht/planning-tools/internal/clients"
	"github.com/derhabicht/planning-tools/pkg/metoc"
)

func ParsePrecipType(raw []string) []metoc.PrecipitationType {
	var types []metoc.PrecipitationType

	for _, t := range raw {
		switch t {
		case "rain":
			types = append(types, metoc.Rain)
		case "snow":
			types = append(types, metoc.Snow)
		case "freezingrain":
			types = append(types, metoc.FreezingRain)
		case "ice":
			types = append(types, metoc.Hail)
		}
	}

	return types
}

func ParseCloudCover(raw float64) metoc.CloudCover {
	if raw < 12.5 {
		return metoc.Clear
	} else if raw < 37.5 {
		return metoc.Scattered
	} else if raw < 62.5 {
		return metoc.Few
	} else if raw < 100.0 {
		return metoc.Broken
	}

	return metoc.Overcast
}

func ParseHourlyForecast(dtg metoc.Dtg, raw clients.VisualCrossingHourlyData) metoc.HourlyForecast {
	return metoc.HourlyForecast{
		DateTime:                 dtg,
		WindDirection:            raw.WindDirection,
		WindSpeed:                raw.WindSpeed,
		WindGust:                 raw.WindGust,
		Visibility:               raw.Visibility,
		PrecipitationAmount:      raw.Precipitation,
		PrecipitationType:        ParsePrecipType(raw.PrecipitationType),
		PrecipitationProbability: raw.PrecipitationProbability,
		CloudCover:               ParseCloudCover(raw.CloudCover),
		Temperature:              raw.Temperature,
		FeelsLike:                raw.FeelsLike,
		Dewpoint:                 raw.Dewpoint,
		Pressure:                 raw.Pressure,
	}
}

func ParseDailyForecast(raw clients.VisualCrossingDailyData, day metoc.Dtg, tzoffset int) (metoc.DailyForecast, error) {
	df := metoc.DailyForecast{
		Conditions:   raw.Conditions,
		Description:  raw.Description,
		HighTemp:     raw.TempMax,
		LowTemp:      raw.TempMin,
		FeelsLikeMax: raw.FeelsLikeMax,
		FeelsLikeMin: raw.FeelsLikeMin,
		Hours:        make(map[int]metoc.HourlyForecast),
	}

	for _, hour := range raw.Hours {
		dt := fmt.Sprintf("%sT%s", day.IsoDate(), hour.Datetime)
		dtg, err := metoc.ParseIsoDateTimeToDtg(dt, tzoffset)
		if err != nil {
			return metoc.DailyForecast{}, errors.WithStack(err)
		}

		hf := ParseHourlyForecast(dtg, hour)
		df.Hours[dtg.Hour()] = hf
	}

	return df, nil
}

func FetchWeatherDataForReport(report *metoc.MetocReport) error {
	data, err := clients.FetchVisualCrossingData(
		viper.GetString("visual_crossing.api_key"),
		report.Location,
		report.Dates[0],
		report.Dates[len(report.Dates)-1],
	)

	if err != nil {
		return errors.WithStack(err)
	}

	for _, day := range data.Days {
		dtg, err := metoc.ParseIsoDateToDtg(day.Datetime)

		df, err := ParseDailyForecast(day, dtg, int(data.TzOffset))
		if err != nil {
			return errors.WithStack(err)
		}

		report.Forecast[dtg] = df
	}

	return nil
}
