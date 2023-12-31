package calendar

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
)

type FyTrimester int

const (
	FyT1 FyTrimester = iota + 1
	FyT2
	FyT3
)

func (f FyTrimester) StartDate(fy int) date.Date {
	switch f {
	case FyT1:
		return date.New(fy-1, time.October, 1)
	case FyT2:
		return date.New(fy, time.February, 1)
	case FyT3:
		return date.New(fy, time.June, 1)
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f FyTrimester) NextTrimester(year int) (int, FyTrimester) {
	switch f {
	case FyT1:
		return year, FyT2
	case FyT2:
		return year, FyT3
	case FyT3:
		return year + 1, FyT1
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f FyTrimester) FullName(fy int) string {
	year := fmt.Sprintf("FY%d", fy)
	switch f {
	case FyT1:
		return fmt.Sprintf("%s, 1st Trimester", year)
	case FyT2:
		return fmt.Sprintf("%s, 2nd Trimester", year)
	case FyT3:
		return fmt.Sprintf("%s, 3rd Trimester", year)
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f FyTrimester) String() string {
	switch f {
	case FyT1:
		return "T1"
	case FyT2:
		return "T2"
	case FyT3:
		return "T3"
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func computeCapTrimester(d date.Date) FyTrimester {
	switch d.Month() {
	case time.October, time.November, time.December, time.January:
		return FyT1
	case time.February, time.March, time.April, time.May:
		return FyT2
	case time.June, time.July, time.August, time.September:
		return FyT3
	default:
		panic(fmt.Errorf("%d is not a valid month", d.Month()))
	}
}

type FyQuarter int

const (
	FyQ1 FyQuarter = iota + 1
	FyQ2
	FyQ3
	FyQ4
)

func (f FyQuarter) CyQuarter(fy int) (int, Ag7ifQuarter) {
	switch f {
	case FyQ1:
		return fy - 1, Ag7ifQ4
	case FyQ2:
		return fy, Ag7ifQ1
	case FyQ3:
		return fy, Ag7ifQ2
	case FyQ4:
		return fy, Ag7ifQ3
	default:
		panic(fmt.Errorf("%d is not a valid quarter", f))
	}
}

func (f FyQuarter) NextQuarter(fy int) (int, FyQuarter) {
	switch f {
	case FyQ1:
		return fy, FyQ2
	case FyQ2:
		return fy, FyQ3
	case FyQ3:
		return fy, FyQ4
	case FyQ4:
		return fy + 1, FyQ1
	default:
		panic(fmt.Errorf("%d is not a valid quarter", f))
	}

}

func (f FyQuarter) StartDate(fy int) date.Date {
	switch f {
	case FyQ1:
		return date.New(fy-1, time.October, 1)
	case FyQ2:
		return date.New(fy, time.January, 1)
	case FyQ3:
		return date.New(fy, time.April, 1)
	case FyQ4:
		return date.New(fy, time.July, 1)
	default:
		panic(fmt.Errorf("%d is not a valid quarter", f))
	}
}

func (f FyQuarter) FullName(fy int) string {
	year := fmt.Sprintf("FY%d", fy)
	switch f {
	case FyQ1:
		return fmt.Sprintf("%s, 1st Quarter", year)
	case FyQ2:
		return fmt.Sprintf("%s, 2nd Quarter", year)
	case FyQ3:
		return fmt.Sprintf("%s, 3rd Quarter", year)
	case FyQ4:
		return fmt.Sprintf("%s, 4th Quarter", year)
	default:
		panic(fmt.Errorf("%d is not a valid trimester", f))
	}
}

func (f FyQuarter) String() string {
	switch f {
	case FyQ1:
		return "Q1"
	case FyQ2:
		return "Q2"
	case FyQ3:
		return "Q3"
	case FyQ4:
		return "Q4"
	default:
		panic(fmt.Errorf("%d is not a valid quarter", f))
	}
}

func computeCapQuarter(d date.Date) FyQuarter {
	switch d.Month() {
	case time.October, time.November, time.December:
		return FyQ1
	case time.January, time.February, time.March:
		return FyQ2
	case time.April, time.May, time.June:
		return FyQ3
	case time.July, time.August, time.September:
		return FyQ4
	default:
		panic(fmt.Errorf("%d is not a valid month", d.Month()))
	}
}

type CapFiscalYear struct {
	year      int
	startDate date.Date
	endDate   date.Date
}

func NewCapFiscalYear(year int) CapFiscalYear {
	sd := date.New(year-1, time.October, 1)
	ed := date.New(year, time.October, 1)

	startDate := computeNearestMonday(sd)
	endDate := computeNearestMonday(ed)

	return CapFiscalYear{
		year:      year,
		startDate: startDate,
		endDate:   endDate,
	}
}

func FiscalYearFromDate(d date.Date) CapFiscalYear {
	switch d.Month() {
	case time.November, time.December:
		return NewCapFiscalYear(d.Year() + 1)
	case time.October:
		fy := NewCapFiscalYear(d.Year() + 1)
		if !fy.IsWithinFiscalYear(d) {
			return NewCapFiscalYear(d.Year())
		}

		return fy
	case time.September:
		fy := NewCapFiscalYear(d.Year())
		if !fy.IsWithinFiscalYear(d) {
			return NewCapFiscalYear(d.Year() + 1)
		}

		return fy
	default:
		return NewCapFiscalYear(d.Year())
	}
}

func (fy CapFiscalYear) IsWithinFiscalYear(d date.Date) bool {
	return d.Equal(fy.startDate) || (d.After(fy.startDate) && d.Before(fy.endDate))
}

func (fy CapFiscalYear) FyWeek(d date.Date) (int, error) {
	if !fy.IsWithinFiscalYear(d) {
		return -1, fmt.Errorf("%s does not fall within FY%d", d.FormatISO(4), fy.year)
	}

	return (d.Sub(fy.startDate) / 7) + 1, nil
}
