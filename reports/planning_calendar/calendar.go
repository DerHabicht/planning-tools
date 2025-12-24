package planning_calendar

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/mod/semver"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/doomsday"
	"github.com/derhabicht/planning-tools/reports/planning_calendar/templates"
)

const calendarMonthCount = 15
const trimesterCount = 4 // 15/4 = 4.75, so we round up
const quarterCount = 5   // 15/3 = 5

type Calendar struct {
	calendar   calendar.Calendar
	minimonths map[string]Minimonth
}

func NewCalendar(cal calendar.Calendar) *Calendar {
	minimonths := NewMinimonthList(cal)

	return &Calendar{
		calendar:   cal,
		minimonths: minimonths,
	}
}

func (c *Calendar) fillCalParams(latex string) string {
	//	Set the Lunar Calibration Date that the Tikz uses to calculate and then draw the phase of the moon
	latex = strings.Replace(latex, templates.LunarCalibrationDate, c.calendar.LunarCalibrationDate().FormatISO(4), 1)
	//	Set the full name and year of the first full month page in this planning_calendar
	latex = strings.Replace(latex, templates.CalendarStart, fmt.Sprintf("October %d", c.calendar.FiscalYear()-1), 1)
	//	Set the full name and year of the last full month page in this planning_calendar
	latex = strings.Replace(latex, templates.CalendarEnd, fmt.Sprintf("December %d", c.calendar.FiscalYear()), 1)
	// Set Natural Year values
	latex = strings.Replace(latex, templates.NaturalYearStart, fmt.Sprintf("%d", c.calendar.NaturalYear()), 1)
	latex = strings.Replace(latex, templates.NaturalYearEnd, fmt.Sprintf("%d", c.calendar.NaturalYear()+1), 1)
	latex = strings.Replace(latex, templates.NaturalYearDecade, fmt.Sprintf("%d", c.calendar.NaturalYearDecade()), 1)
	latex = strings.Replace(latex, templates.NaturalYearLustrum, fmt.Sprintf("%d", c.calendar.NaturalYearLustrum()), 1)
	latex = strings.Replace(latex, templates.NaturalYearTriad, fmt.Sprintf("%d", c.calendar.NaturalYearTriad()), 1)

	//	Set the starting year of this planning_calendar, expressed as the year of Julian Period A
	latex = strings.Replace(latex, templates.JulianPeriodStart, strconv.Itoa(c.calendar.JulianPeriod()-1), 1)
	//	Set the ending year of this planning_calendar, expressed as the year of Julian Period A
	latex = strings.Replace(latex, templates.JulianPeriodEnd, strconv.Itoa(c.calendar.JulianPeriod()), 1)
	//	Set the picture to typeset on the title page of the planning_calendar
	latex = strings.Replace(latex, templates.TitlePicture, config.GetString(config.CoverLogo), 1)
	//	Set the first planning_calendar year covered in this planning_calendar (i.e. FY-1)
	latex = strings.Replace(latex, templates.CalendarYear1, strconv.Itoa(c.calendar.FiscalYear()-1), -1)
	//	Set the second planning_calendar year covered in this planning_calendar (i.e. FY)
	latex = strings.Replace(latex, templates.CalendarYear2, strconv.Itoa(c.calendar.FiscalYear()), -1)
	//	Set the current version of plancal
	latex = strings.Replace(latex, templates.PlanCalV, semver.Canonical(fmt.Sprintf("v%s", config.GetString(config.Version))), -1)

	//	Set the color to use for the title box outline on the planning_calendar's title page
	if c.calendar.FiscalYear()%2 == 0 {
		latex = strings.Replace(latex, templates.TitleColor, "blue", 1)
	} else {
		latex = strings.Replace(latex, templates.TitleColor, "red", 1)
	}

	return latex
}

func (c *Calendar) generateAG7IFTables(latex string) string {
	const layout = `02 Jan`

	table1 := templates.AG7IFTable
	table2 := templates.AG7IFTable

	cy1week, _ := c.calendar.FetchWeek(c.calendar.FiscalYear()-1, 1)
	cy2week, _ := c.calendar.FetchWeek(c.calendar.FiscalYear(), 1)

	for i := 1; i <= 52; i++ {
		cy1wkst := cy1week.StartDay().Date().Format(layout)
		cy1wkend := cy1week.StartDay().Date().Add(6).Format(layout)

		cy2wkst := cy2week.StartDay().Date().Format(layout)
		cy2wkend := cy2week.StartDay().Date().Add(6).Format(layout)

		table1 = strings.Replace(table1, templates.Ag7IFTableWeekStart(i), cy1wkst, 1)
		table1 = strings.Replace(table1, templates.AG7IFTableWeekEnd(i), cy1wkend, 1)

		table2 = strings.Replace(table2, templates.Ag7IFTableWeekStart(i), cy2wkst, 1)
		table2 = strings.Replace(table2, templates.AG7IFTableWeekEnd(i), cy2wkend, 1)

		cy1week = cy1week.Next()
		cy2week = cy2week.Next()
	}

	latex = strings.Replace(latex, templates.AG7IFTable1, table1, 1)
	latex = strings.Replace(latex, templates.AG7IFTable2, table2, 1)

	return latex
}

func (c *Calendar) generateDoomsdayTable(latex string) string {
	table := templates.DoomsdayTableTemplate

	var rows string
	for year := c.calendar.FiscalYear() - 2; year <= c.calendar.FiscalYear()+2; year++ {
		dd := doomsday.ComputeDoomsday(year)
		row := templates.DoomsdayTableRowTemplate
		if year == c.calendar.FiscalYear() {
			row = strings.Replace(row, templates.Year, fmt.Sprintf(`\textbf{%d}`, year), 1)
			row = strings.Replace(row, templates.Doomsday, fmt.Sprintf(`\textbf{%s}`, calendar.WeekdayLetter(dd)), 1)
		} else {
			row = strings.Replace(row, templates.Year, strconv.Itoa(year), 1)
			row = strings.Replace(row, templates.Doomsday, calendar.WeekdayLetter(dd), 1)
		}

		rows += row
	}

	table = strings.Replace(table, templates.DoomsdayTableRows, rows, 1)
	latex = strings.Replace(latex, templates.DoomsdayTable, table, 1)

	return latex
}

func (c *Calendar) generateSolsticeTable(latex string) string {
	const layout = "021504Z Jan"

	table := templates.SolsticeTableTemplate
	table = strings.Replace(table, templates.CalendarYear1, strconv.Itoa(c.calendar.FiscalYear()-1), 1)
	table = strings.Replace(table, templates.WinterSolstice1, c.calendar.SolsticeTable().FirstWinterSolstice().Format(layout), 1)
	table = strings.Replace(table, templates.VernalEquinox, c.calendar.SolsticeTable().VernalEquinox().Format(layout), 1)
	table = strings.Replace(table, templates.SummerSolstice, c.calendar.SolsticeTable().SummerSolstice().Format(layout), 1)
	table = strings.Replace(table, templates.AutumnalEquinox, c.calendar.SolsticeTable().AutumnalEquinox().Format(layout), 1)
	table = strings.Replace(table, templates.CalendarYear2, strconv.Itoa(c.calendar.FiscalYear()), 1)
	table = strings.Replace(table, templates.WinterSolstice2, c.calendar.SolsticeTable().SecondWinterSolstice().Format(layout), 1)

	latex = strings.Replace(latex, templates.SolsticeTable, table, 1)

	return latex
}

func (c *Calendar) generateHolidayTables(latex string) string {
	ht := NewHolidayTables(c.calendar.HolidayCalendar(), c.calendar.FiscalYear())

	latex = ht.TableByOccurrence(latex)
	latex = ht.TableByAbbreviation(latex)

	return latex
}

func (c *Calendar) generateMiniMonthCmds(latex string) string {
	mm := ""
	for _, m := range c.minimonths {
		mm += m.LaTeX()
	}

	latex = strings.Replace(latex, templates.MinimonthCommands, mm, 1)

	return latex
}

func (c *Calendar) generateTrimesterPages(latex string) string {
	trimester := c.calendar.FirstTrimester()
	for i := 1; i <= trimesterCount; i++ {
		tr := NewTrimester(trimester, c.minimonths)

		latex = strings.Replace(latex, templates.TrimesterPage(i), tr.LaTeX(), 1)

		trimester = trimester.Next()
	}

	return latex
}

func (c *Calendar) generateQuarterPages(latex string) string {
	calQtr := c.calendar.FirstCalendarQuarter()
	fyQtr := c.calendar.FirstFiscalQuarter()
	for i := 1; i <= quarterCount; i++ {
		qt := NewQuarter(calQtr, fyQtr, c.minimonths)

		latex = strings.Replace(latex, templates.QuarterPage(i), qt.LaTeX(), 1)

		calQtr = calQtr.Next()
		fyQtr = fyQtr.Next()
	}

	return latex
}

func (c *Calendar) generateMonthPages(latex string) string {
	month := c.calendar.FirstMonth()
	for i := 1; i <= calendarMonthCount; i++ {
		mo := NewMonth(c.calendar, month, c.minimonths)
		latex = strings.Replace(latex, templates.MonthPage(i), mo.LaTeX(), 1)
		month = month.Next()
	}

	return latex
}

func (c *Calendar) LaTeX() string {
	latex := templates.CalendarTemplate

	latex = c.fillCalParams(latex)
	latex = c.generateDoomsdayTable(latex)
	latex = c.generateSolsticeTable(latex)
	latex = c.generateAG7IFTables(latex)
	latex = c.generateHolidayTables(latex)
	latex = c.generateMiniMonthCmds(latex)
	latex = c.generateTrimesterPages(latex)
	latex = c.generateQuarterPages(latex)
	latex = c.generateMonthPages(latex)

	return latex
}
