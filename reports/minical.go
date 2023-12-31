package reports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fxtlabs/date"
)

const miniCalWeekTemplate = `+W & +D1 & +D2 & +D3 & +D4 & +D5 & +D6 & +D7 \\`

func fillMiniCalMonthNames(template string, startDate date.Date, months int) string {
	d := startDate

	for i := 0; i < months; i++ {
		template = strings.Replace(template, fmt.Sprintf("+M%dTITLE", i+1), d.Format("January 2006"), 1)
		d = d.AddDate(0, 1, 0)
	}

	return template
}

func fillMiniCalWeeks(template string, start date.Date, month int) string {
	d := start
	startMonth := d.Month()
	for week := 1; week <= 6; week++ {
		wkTemplate := miniCalWeekTemplate
		if d.Month() == startMonth {
			_, isoWeek := d.ISOWeek()
			wkTemplate = strings.Replace(wkTemplate, "+W", strconv.Itoa(isoWeek), 1)
		} else {
			wkTemplate = strings.Replace(wkTemplate, "+W", "", 1)
		}
		for weekDay := 1; weekDay <= 7; weekDay++ {
			if (int(d.Weekday()) != weekDay%7) || d.Month() != startMonth {
				wkTemplate = strings.Replace(wkTemplate, fmt.Sprintf("+D%d", weekDay), "", 1)
				continue
			}
			wkTemplate = strings.Replace(wkTemplate, fmt.Sprintf("+D%d", weekDay), strconv.Itoa(d.Day()), 1)
			d = d.Add(1)
		}

		template = strings.Replace(template, fmt.Sprintf("+M%dW%d", month, week), wkTemplate, 1)
	}

	return template
}

func fillMiniCalMonths(template string, startDate date.Date, months int) string {
	d := startDate
	for i := 0; i < months; i++ {
		template = fillMiniCalWeeks(template, d, i+1)
		d = d.AddDate(0, 1, 0)
	}

	return template
}
