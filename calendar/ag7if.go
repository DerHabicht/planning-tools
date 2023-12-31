package calendar

import (
	"fmt"
	"time"
)

type Suit int

const (
	Hearts Suit = iota + 1
	Clubs
	Diamonds
	Spades
)

// Card represents a playing card.
// Rank values have the following meanings:
//
//		0: Joker
//	 1: Ace
//		2: 2
//	 ...
//	 10: 10
//	 11: Jack
//	 12: Queen
//	 13: King
type Card struct {
	Suit Suit
	Rank uint
}

var weekCard map[int]Card = map[int]Card{
	1:  Card{Suit: Hearts, Rank: 1},
	2:  Card{Suit: Hearts, Rank: 2},
	3:  Card{Suit: Hearts, Rank: 3},
	4:  Card{Suit: Hearts, Rank: 4},
	5:  Card{Suit: Hearts, Rank: 5},
	6:  Card{Suit: Hearts, Rank: 6},
	7:  Card{Suit: Hearts, Rank: 7},
	8:  Card{Suit: Hearts, Rank: 8},
	9:  Card{Suit: Hearts, Rank: 9},
	10: Card{Suit: Hearts, Rank: 10},
	11: Card{Suit: Hearts, Rank: 11},
	12: Card{Suit: Hearts, Rank: 12},
	13: Card{Suit: Hearts, Rank: 13},

	14: Card{Suit: Clubs, Rank: 1},
	15: Card{Suit: Clubs, Rank: 2},
	16: Card{Suit: Clubs, Rank: 3},
	17: Card{Suit: Clubs, Rank: 4},
	18: Card{Suit: Clubs, Rank: 5},
	19: Card{Suit: Clubs, Rank: 6},
	20: Card{Suit: Clubs, Rank: 7},
	21: Card{Suit: Clubs, Rank: 8},
	22: Card{Suit: Clubs, Rank: 9},
	23: Card{Suit: Clubs, Rank: 10},
	24: Card{Suit: Clubs, Rank: 11},
	25: Card{Suit: Clubs, Rank: 12},
	26: Card{Suit: Clubs, Rank: 13},

	27: Card{Suit: Diamonds, Rank: 13},
	28: Card{Suit: Diamonds, Rank: 12},
	29: Card{Suit: Diamonds, Rank: 11},
	30: Card{Suit: Diamonds, Rank: 10},
	31: Card{Suit: Diamonds, Rank: 9},
	32: Card{Suit: Diamonds, Rank: 8},
	33: Card{Suit: Diamonds, Rank: 7},
	34: Card{Suit: Diamonds, Rank: 6},
	35: Card{Suit: Diamonds, Rank: 5},
	36: Card{Suit: Diamonds, Rank: 4},
	37: Card{Suit: Diamonds, Rank: 3},
	38: Card{Suit: Diamonds, Rank: 2},
	39: Card{Suit: Diamonds, Rank: 1},

	40: Card{Suit: Spades, Rank: 13},
	41: Card{Suit: Spades, Rank: 12},
	42: Card{Suit: Spades, Rank: 11},
	43: Card{Suit: Spades, Rank: 10},
	44: Card{Suit: Spades, Rank: 9},
	45: Card{Suit: Spades, Rank: 8},
	46: Card{Suit: Spades, Rank: 7},
	47: Card{Suit: Spades, Rank: 6},
	48: Card{Suit: Spades, Rank: 5},
	49: Card{Suit: Spades, Rank: 4},
	50: Card{Suit: Spades, Rank: 3},
	51: Card{Suit: Spades, Rank: 2},
	52: Card{Suit: Spades, Rank: 1},

	53: Card{Suit: 0, Rank: 0},
}

func GetWeekCard(isoWeek int) Card {
	return weekCard[isoWeek]
}

type Ag7ifSprint int

const (
	S01 Ag7ifSprint = iota + 1
	S02
	S03
	S04
	S05
	S06
	S07
	S08
	S09
	S10
	S11
	S12
	S13
	S14
	S15
	S16
	S17
	S18
	S19
	S20
	S21
	S22
	S23
	S24
	SP1
	SP2
	SP3
	SP4
)

func computeAg7ifSprint(isoWeek int) Ag7ifSprint {
	switch isoWeek {
	case 1, 2:
		return S01
	case 3, 4:
		return S02
	case 5, 6:
		return S03
	case 7, 8:
		return S04
	case 9, 10:
		return S05
	case 11, 12:
		return S06
	case 13:
		return SP1
	case 14, 15:
		return S07
	case 16, 17:
		return S08
	case 18, 19:
		return S09
	case 20, 21:
		return S10
	case 22, 23:
		return S11
	case 24, 25:
		return S12
	case 26:
		return SP3
	case 27, 28:
		return S13
	case 29, 30:
		return S14
	case 31, 32:
		return S15
	case 33, 34:
		return S16
	case 35, 36:
		return S17
	case 37, 38:
		return S18
	case 39:
		return SP3
	case 40, 41:
		return S19
	case 42, 43:
		return S20
	case 44, 45:
		return S21
	case 46, 47:
		return S22
	case 48, 49:
		return S23
	case 50, 51:
		return S24
	case 52, 53:
		return SP4
	default:
		panic(fmt.Errorf("%d is not a valid week number", isoWeek))
	}
}

func (a Ag7ifSprint) String() string {
	switch a {
	case S01:
		return "S01"
	case S02:
		return "S02"
	case S03:
		return "S03"
	case S04:
		return "S04"
	case S05:
		return "S05"
	case S06:
		return "S06"
	case S07:
		return "S07"
	case S08:
		return "S08"
	case S09:
		return "S09"
	case S10:
		return "S10"
	case S11:
		return "S11"
	case S12:
		return "S12"
	case S13:
		return "S13"
	case S14:
		return "S14"
	case S15:
		return "S15"
	case S16:
		return "S16"
	case S17:
		return "S17"
	case S18:
		return "S18"
	case S19:
		return "S19"
	case S20:
		return "S20"
	case S21:
		return "S21"
	case S22:
		return "S22"
	case S23:
		return "S23"
	case S24:
		return "S24"
	case SP1:
		return "SP1"
	case SP2:
		return "SP2"
	case SP3:
		return "SP3"
	case SP4:
		return "SP4"
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF sprint", a))
	}
}

type Ag7ifQuarter int

const (
	Ag7ifQ1 Ag7ifQuarter = iota + 1
	Ag7ifQ2
	Ag7ifQ3
	Ag7ifQ4
)

func computeAg7ifQuarter(isoWeek int) Ag7ifQuarter {
	switch {
	case (1 <= isoWeek) && (isoWeek <= 13):
		return Ag7ifQ1
	case (14 <= isoWeek) && (isoWeek <= 26):
		return Ag7ifQ2
	case (27 <= isoWeek) && (isoWeek <= 39):
		return Ag7ifQ3
	case (40 <= isoWeek) && (isoWeek <= 53):
		return Ag7ifQ4
	default:
		panic(fmt.Errorf("%d is not a valid week number", isoWeek))
	}
}

func (a Ag7ifQuarter) StartWeek() int {
	switch a {
	case Ag7ifQ1:
		return 1
	case Ag7ifQ2:
		return 14
	case Ag7ifQ3:
		return 27
	case Ag7ifQ4:
		return 40
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF quarter", a))
	}
}

func (a Ag7ifQuarter) StartMonth() time.Month {
	switch a {
	case Ag7ifQ1:
		return time.January
	case Ag7ifQ2:
		return time.April
	case Ag7ifQ3:
		return time.June
	case Ag7ifQ4:
		return time.October
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF quarter", a))
	}
}

func (a Ag7ifQuarter) FullName(fy int) string {
	year := fmt.Sprintf("CY%d", fy)
	switch a {
	case Ag7ifQ1:
		return fmt.Sprintf("%s, 1st Quarter", year)
	case Ag7ifQ2:
		return fmt.Sprintf("%s, 2nd Quarter", year)
	case Ag7ifQ3:
		return fmt.Sprintf("%s, 3rd Quarter", year)
	case Ag7ifQ4:
		return fmt.Sprintf("%s, 4th Quarter", year)
	default:
		panic(fmt.Errorf("%d is not a valid trimester", a))
	}
}

func (a Ag7ifQuarter) String() string {
	switch a {
	case Ag7ifQ1:
		return "Q1"
	case Ag7ifQ2:
		return "Q2"
	case Ag7ifQ3:
		return "Q3"
	case Ag7ifQ4:
		return "Q4"
	default:
		panic(fmt.Errorf("%d is not a valid AG7IF quarter", a))
	}
}
