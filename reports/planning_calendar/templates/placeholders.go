package templates

import (
	"fmt"
)

const (
	AutumnalEquinox            = `+AE`
	CalendarEnd                = `+CAL_END`
	CalendarQuarter            = `+AQ`
	CalendarStart              = `+CAL_START`
	CalendarYear1              = `+CY1`
	CalendarYear2              = `+CY2`
	Doomsday                   = `+DD`
	DoomsdayTable              = `+DOOMSDAYS`
	DoomsdayTableRows          = `+DD_TABLE_ROWS`
	FiscalQuarter              = `+FQ`
	FiscalTrimester            = `+FT`
	FiscalWeek                 = `+FW`
	FullCalendarQuarter        = `+CYQ`
	FullDate                   = `+FD`
	FullFiscalQuarter          = `+FYQ`
	FullTrimester              = `+T`
	Holiday                    = `+HD`
	HolidayAbbreviation        = `+ABBV`
	HolidayActual1             = `+CY1Act`
	HolidayActual2             = `+CY2Act`
	HolidayName                = `+FULL_NAME`
	HolidayObserved1           = `+CY1Obs`
	HolidayObserved2           = `+CY2Obs`
	HolidayTableByAbbreviation = `+ABBVS`
	HolidayTableByOccurrence   = `+HOLIDAYS`
	ISOWeek                    = `+IW`
	JulianPeriodEnd            = `+JP_END`
	JulianPeriodStart          = `+JP_START`
	LunarCalibrationDate       = `+LCD`
	MJD                        = `+MJD`
	MinimonthCommand           = `+COMMAND`
	MinimonthCommands          = `+MINIMONTH_CMDS`
	MinimonthNext              = `+NEXT_CMD`
	MinimonthPrevious          = `+PREV_CMD`
	MinimonthWeekNumber        = `+W`
	MonthDay                   = `+DY`
	MonthName                  = `+MONTH`
	MonthNameFull              = `+M`
	OKRHeader                  = `+OKR_HDR`
	OrdinalDay                 = `+YD`
	PlanCalV                   = `+PLANCALV`
	SolsticeTable              = `+SOLSTICES`
	Sprint                     = `+AS`
	SummerSolstice             = `+SS`
	Sunday                     = `+SUN`
	SunriseTime                = `+SR`
	SunsetTime                 = `+SS`
	TitleColor                 = `+TITLE_COLOR`
	TitlePicture               = `+PIC`
	VernalEquinox              = `+VE`
	WeekData                   = `+WK`
	WeekHeader                 = `+WEEK_HEADER`
	WeekdayHeader              = `+WEEKDAYS`
	WinterSolstice1            = `+WS1`
	WinterSolstice2            = `+WS2`
	Year                       = `+Y`
	minimonthDay               = `+D%d`
	minimonthMacro             = `+M%dCMD`
	minimonthWeek              = `+W%d`
	monthDayData               = `+D%d`
	monthPage                  = `+M%02d`
	monthPageWeek              = `+W%d`
	okrHeaderWeekNumber        = `+W%02d`
	quarterPage                = `+Q%d`
	trimesterPage              = `+T%d`
)

func MinimonthDay(day int) string {
	return fmt.Sprintf(minimonthDay, day)
}

func MinimonthMacro(month int) string {
	return fmt.Sprintf(minimonthMacro, month)
}

func MinimonthWeek(week int) string {
	return fmt.Sprintf(minimonthWeek, week)
}

func MonthDayData(day int) string {
	return fmt.Sprintf(monthDayData, day)
}

func MonthPage(page int) string {
	return fmt.Sprintf(monthPage, page)
}

func MonthPageWeek(week int) string {
	return fmt.Sprintf(monthPageWeek, week)
}

func OKRHeaderWeekNumber(week int) string {
	return fmt.Sprintf(okrHeaderWeekNumber, week)
}

func QuarterPage(page int) string {
	return fmt.Sprintf(quarterPage, page)
}

func TrimesterPage(page int) string {
	return fmt.Sprintf(trimesterPage, page)
}
