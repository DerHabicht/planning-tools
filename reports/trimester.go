package reports

import (
	"fmt"
	"strings"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/reports/templates"
)

const trimesterMonthCount = 4

type Trimester struct {
	trimester  calendar.Trimester
	minimonths []Minimonth
}

func NewTrimester(trimester calendar.Trimester, minimonths map[string]Minimonth) Trimester {
	var mm []Minimonth

	mo := trimester.FirstMonth()
	for i := 0; i < trimesterMonthCount; i++ {
		mm = append(mm, minimonths[mo.Full()])
		mo = mo.Next()
	}

	return Trimester{
		trimester:  trimester,
		minimonths: mm,
	}
}

func (t *Trimester) LaTeX() string {
	latex := templates.TrimesterTemplate

	latex = strings.Replace(latex, templates.FullTrimester, t.trimester.Full(), 1)

	for i, mm := range t.minimonths {
		latex = strings.Replace(latex, fmt.Sprintf(templates.MinimonthMacro, i+1), mm.LatexCommand(), 1)
	}

	return latex
}
