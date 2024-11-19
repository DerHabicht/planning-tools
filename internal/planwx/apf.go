package planwx

import (
	"fmt"
	"sort"

	"github.com/derhabicht/planning-tools/pkg/metoc"
)

func EncodeRemarks(hf metoc.HourlyForecast) string {
	rmk := "RMK "

	heatCat := metoc.CalculateHeatCategory(hf.FeelsLike)
	if heatCat != metoc.NoRisk {
		rmk += fmt.Sprintf("HEAT RISK %s ", heatCat)
	}

	coldCat := metoc.CalculateColdCategory(hf.FeelsLike)
	if coldCat != metoc.NoRisk {
		rmk += fmt.Sprintf("COLD RISK %s ", coldCat)
	}

	if hf.PrecipitationProbability >= 10 {
		rmk += fmt.Sprintf("PRECIP PROB %.0f%%", hf.PrecipitationProbability)
	}

	if rmk != "RMK " {
		return rmk
	}

	return ""
}

func EncodePressure(pressure float64) string {
	inhg := pressure / 33.864

	return fmt.Sprintf("A%04.0f", inhg)
}

func EncodeTemperature(temp float64, dewpoint float64) string {
	tempStr := ""
	dewStr := ""

	if temp < 0 {
		temp *= -1
		tempStr = fmt.Sprintf("M%02.0f", temp)
	} else {
		tempStr = fmt.Sprintf("%02.0f", temp)
	}

	if dewpoint < 0 {
		dewpoint *= -1
		dewStr = fmt.Sprintf("M%02.0f", dewpoint)
	} else {
		dewStr = fmt.Sprintf("%02.0f", dewpoint)
	}

	return fmt.Sprintf("%s/%s", tempStr, dewStr)
}

func EncodeWeather(precipType []metoc.PrecipitationType, precipAmount float64) string {
	if len(precipType) < 1 {
		return ""
	}

	intensity := metoc.CalculatePrecipitationIntensity(precipAmount)

	prefix := ""
	switch intensity {
	case metoc.HeavyPrecip:
		prefix = "+"
	case metoc.LightPrecip:
		prefix = "-"
	case metoc.NoPrecip:
		return ""
	}

	sort.Slice(precipType, func(i, j int) bool { return precipType[i] < precipType[j] })

	wx := ""
	switch precipType[0] {
	case metoc.FreezingRain:
		wx = "FZRA"
	case metoc.Hail:
		wx = "GR"
	case metoc.Snow:
		wx = "SN"
	case metoc.Rain:
		wx = "RN"
	}

	return fmt.Sprintf("%s%s", prefix, wx)
}

func EncodeVisibility(visibility float64) string {
	return fmt.Sprintf("%.0fSM", visibility/1.609)
}

func EncodeWind(direction, speed, gust float64) string {
	speed_knots := speed / 1.852
	gust_knots := gust / 1.852

	if (gust_knots - speed_knots) < 5 {
		return fmt.Sprintf("%03.0f%02.f", direction, speed_knots)
	}

	return fmt.Sprintf("%03.0f%02.0fG%02.0f", direction, speed_knots, gust_knots)
}

func EncodeHourlyForecast(hf metoc.HourlyForecast) string {
	fcst := fmt.Sprintf("FM%s", hf.DateTime.Short())
	wind := EncodeWind(hf.WindDirection, hf.WindSpeed, hf.WindGust)
	visibility := EncodeVisibility(hf.Visibility)
	wx := EncodeWeather(hf.PrecipitationType, hf.PrecipitationAmount)
	sky := string(hf.CloudCover)
	temp := EncodeTemperature(hf.Temperature, hf.Dewpoint)
	pressure := EncodePressure(hf.Pressure)
	remarks := EncodeRemarks(hf)

	if wx != "" {
		return fmt.Sprintf("%s %s %s %s %s %s %s %s\n",
			fcst,
			wind,
			visibility,
			wx,
			sky,
			temp,
			pressure,
			remarks,
		)
	}

	return fmt.Sprintf("%s %s %s %s %s %s %s\n",
		fcst,
		wind,
		visibility,
		sky,
		temp,
		pressure,
		remarks,
	)
}

func EncodeDailyForecast(df metoc.DailyForecast, location metoc.Location, generated metoc.Dtg) string {
	fcst := fmt.Sprintf("APF %s %s\n", location.Mgrs(), generated.Short())

	for i := 0; i < 24; i++ {
		fcst += EncodeHourlyForecast(df.Hours[i])
	}

	return fcst
}
