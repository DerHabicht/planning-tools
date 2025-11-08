package kfm

import (
	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-tools/pkg/calendar"
)

type KFM struct {
	FY       int            `yaml:"fy"`
	Subjects map[int]string `yaml:"subjects"`
	Lessons  []Lesson       `yaml:"lessons"`
}

type Lesson struct {
	Week      int       `yaml:"week"`
	StartDate date.Date `yaml:"start_date"`
	EndDate   date.Date `yaml:"end_date"`
	Topic     string    `yaml:"topic"`
	Reading   []Reading `yaml:"reading"`
}

type Reading struct {
	Book     Book  `yaml:"book"`
	Chapters []int `yaml:"chapters"`
}

func GenerateBlankKFM(cal calendar.Calendar) KFM {
	totalWeeks := calendar.ISOWeekCount(cal.FiscalYear()-1) + calendar.ISOWeekCount(cal.FiscalYear())
	wk := cal.FirstWeek()
	_, wn, _ := wk.ISOWeek()

	totalWeeks -= wn - 1
	var lessons []Lesson
	for i := 1; i <= totalWeeks; i++ {
		_, wn, _ = wk.ISOWeek()
		lesson := Lesson{
			Week:      wn,
			StartDate: wk.StartDay().Date(),
			EndDate:   wk.StartDay().Date().Add(6),
			Topic:     "",
			Reading: []Reading{
				{
					Book:     Genesis,
					Chapters: []int{0},
				},
			},
		}

		lessons = append(lessons, lesson)
		wk = wk.Next()
	}

	return KFM{
		FY: cal.FiscalYear(),
		Subjects: map[int]string{
			cal.FiscalYear() - 1: "",
			cal.FiscalYear():     "",
		},
		Lessons: lessons,
	}
}
