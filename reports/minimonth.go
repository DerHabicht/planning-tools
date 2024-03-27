package reports

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-calendar/calendar"
)

const miniCalWeekHeaderTemplate = `W & +M & +T & +W & +H & +F & +S & +U \\`

const miniCalMonthTemplate = `\fbox{\begin{minipage}{0.24\textwidth}
          \centering
          {\Large\textbf{+MONTH}}\vspace{\baselineskip}
          \begin{tabularx}{\textwidth}{r|rrrrrrr}
              \toprule
 			  +WEEKHEADER
              \midrule
              +W1
              +W2
              +W3
              +W4
              +W5
              +W6
          \end{tabularx}
\end{minipage}}
`
const miniCalWeekTemplate = `+W & +D1 & +D2 & +D3 & +D4 & +D5 & +D6 & +D7 \\`

func isSecondMonth(fy int, d date.Date) bool {
	prevYearDoubledMonth := map[time.Month]bool{
		time.September: true,
		time.October:   true,
		time.November:  true,
		time.December:  true,
	}

	if d.Year() == fy {
		_, ok := prevYearDoubledMonth[d.Month()]
		return ok
	}

	if d.Year() == fy+1 {
		return d.Month() == time.January
	}

	return false
}

func generateTemplates(fy int) []miniMonthTemplate {
	d := date.New(fy-1, time.September, 1)
	endDate := date.New(fy+1, time.February, 1)

	var templates []miniMonthTemplate
	for d.Before(endDate) {
		second := isSecondMonth(fy, d)
		t := newMiniMonthTemplate(d.Year(), d.Month(), second)

		templates = append(templates, t)

		d = d.AddDate(0, 1, 0)
	}

	return templates
}

type miniMonthTemplate struct {
	month          time.Month
	year           int
	doomsday       time.Weekday
	latexCommand   string
	calTemplateKey string
}

func newMiniMonthTemplate(year int, month time.Month, second bool) miniMonthTemplate {
	doubledMonth := map[time.Month]bool{
		time.January:   true,
		time.September: true,
		time.October:   true,
		time.November:  true,
		time.December:  true,
	}

	latexCommand := `\mini`
	calTemplateKey := `+MM`
	if _, ok := doubledMonth[month]; ok {
		if second {
			latexCommand = fmt.Sprintf("%ssecond%s", latexCommand, strings.ToLower(month.String()))
			calTemplateKey = fmt.Sprintf("%s%s2", calTemplateKey, strings.ToUpper(month.String()[:3]))
		} else {
			latexCommand = fmt.Sprintf("%sfirst%s", latexCommand, strings.ToLower(month.String()))
			calTemplateKey = fmt.Sprintf("%s%s1", calTemplateKey, strings.ToUpper(month.String()[:3]))
		}
	} else {
		latexCommand = fmt.Sprintf("%s%s", latexCommand, strings.ToLower(month.String()))
		calTemplateKey = fmt.Sprintf("%s%s", calTemplateKey, strings.ToUpper(month.String()[:3]))
	}

	return miniMonthTemplate{
		month:          month,
		year:           year,
		doomsday:       calendar.ComputeDoomsday(year),
		latexCommand:   latexCommand,
		calTemplateKey: calTemplateKey,
	}
}

func (mmt *miniMonthTemplate) generateWeekHeader() string {
	header := miniCalWeekHeaderTemplate
	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		weekdayLetter := calendar.WeekdayLetters[weekday]

		repl := weekdayLetter
		if weekday == mmt.doomsday {
			repl = fmt.Sprintf("\\underline{%s}", weekdayLetter)
		}

		header = strings.Replace(header, fmt.Sprintf("+%s", weekdayLetter), repl, 1)
	}

	return header
}

func (mmt *miniMonthTemplate) fillWeeks(template string, start date.Date) string {
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

		template = strings.Replace(template, fmt.Sprintf("+W%d", week), wkTemplate, 1)
	}

	return template
}

func (mmt *miniMonthTemplate) MonthStr() string {
	return fmt.Sprintf("%s %d", mmt.month, mmt.year)
}

func (mmt *miniMonthTemplate) TemplateKey() string {
	return mmt.calTemplateKey
}

func (mmt *miniMonthTemplate) LaTeXCommand() string {
	return mmt.latexCommand
}

func (mmt *miniMonthTemplate) LaTeX() string {
	latex := miniCalMonthTemplate
	latex = strings.Replace(latex, "+MONTH", fmt.Sprintf("%s %d", mmt.month, mmt.year), 1)
	latex = strings.Replace(latex, "+WEEKHEADER", mmt.generateWeekHeader(), 1)
	latex = mmt.fillWeeks(latex, date.New(mmt.year, mmt.month, 1))

	return latex
}
