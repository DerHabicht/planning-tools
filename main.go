package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/config"
	"github.com/derhabicht/planning-calendar/latex"
)

func main() {
	config.InitConfig()

	fiscalYear, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("%s is not a valid fiscal year", os.Args[1]))
	}
	outputFile := os.Args[2]

	cal := calendar.NewCalendar(fiscalYear)

	template, err := latex.NewCalendarTemplate(cal)
	if err != nil {
		panic(err)
	}

	generated := []byte(template.LaTeX())
	err = os.WriteFile(outputFile, generated, 0644)
	if err != nil {
		panic(err)
	}
}
