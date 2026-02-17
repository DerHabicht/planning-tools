package planning_calendar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"
)

var tabValues = []string{
	templates.TrimesterPage(1),
	templates.QuarterPage(1),
	templates.MonthPage(1),
	templates.MonthPage(2),
	templates.MonthPage(3),
	templates.QuarterPage(2),
	templates.MonthPage(4),
	templates.TrimesterPage(2),
	templates.MonthPage(5),
	templates.MonthPage(6),
	templates.QuarterPage(3),
	templates.MonthPage(7),
	templates.MonthPage(8),
	templates.TrimesterPage(3),
	templates.MonthPage(9),
	templates.QuarterPage(4),
	templates.MonthPage(10),
	templates.MonthPage(11),
	templates.MonthPage(12),
	templates.TrimesterPage(4),
	templates.QuarterPage(5),
	templates.MonthPage(13),
	templates.MonthPage(14),
	templates.MonthPage(15),
	"Reference Tables",
	"Holidays",
	fmt.Sprintf("%s POI 50-1", templates.CalendarYear1),
	fmt.Sprintf("%s POI 50-1", templates.CalendarYear2),
	"Extended Calendars",
}

type CalendarTabs struct {
	calendar calendar.Calendar
	startTab int
}

func NewCalendarTabs(calendar calendar.Calendar, startTab int) *CalendarTabs {
	return &CalendarTabs{calendar: calendar, startTab: startTab}
}

func (c *CalendarTabs) GeneratePages() string {
	pages := templates.CalendarTabsPage

	for i := 1; i < c.startTab; i++ {
		pages = strings.Replace(pages, templates.TabOrd(i), "", 1)
	}

	tabOrd := c.startTab
	for _, p := range tabValues {
		pages = strings.Replace(pages, templates.TabOrd(tabOrd), p, 1)

		tabOrd++
		if tabOrd > 20 {
			tabOrd = 1
			pages = pages + "\n\n" + `\newpage` + "\n\n" + templates.CalendarTabsPage
		}
	}

	for i := tabOrd; i <= 20; i++ {
		pages = strings.Replace(pages, templates.TabOrd(i), "", 1)
	}

	return pages
}

func (c *CalendarTabs) FillPages(pages string) string {
	pages = strings.Replace(
		pages,
		templates.CalendarYear1,
		strconv.Itoa(c.calendar.FirstMonth().Year()),
		1,
	)
	pages = strings.Replace(
		pages,
		templates.CalendarYear2,
		strconv.Itoa(c.calendar.FiscalYear()),
		1,
	)

	tri := c.calendar.FirstTrimester()
	for i := 1; i <= trimesterCount; i++ {
		pages = strings.Replace(pages, templates.TrimesterPage(i), tri.String(), 1)
		tri = tri.Next()
	}

	cqtr := c.calendar.FirstCalendarQuarter()
	fqtr := c.calendar.FirstFiscalQuarter()
	for i := 1; i <= quarterCount; i++ {
		qtr := fmt.Sprintf(`%s\\%s`, cqtr.String(), fqtr.String())
		pages = strings.Replace(pages, templates.QuarterPage(i), qtr, 1)
		cqtr = cqtr.Next()
		fqtr = fqtr.Next()
	}

	mo := c.calendar.FirstMonth()
	for i := 1; i <= calendarMonthCount; i++ {
		pages = strings.Replace(pages, templates.MonthPage(i), mo.Full(), 1)
		mo = mo.Next()
	}

	return pages
}

func (c *CalendarTabs) LaTeX() string {
	pages := c.GeneratePages()
	pages = c.FillPages(pages)

	return strings.Replace(templates.CalendarTabsTemplate, templates.TabPages, pages, 1)
}
