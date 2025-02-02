package plancal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalendar_FetchWeek(t *testing.T) {
	cal := NewCalendar(2025)

	wk, err := cal.FetchWeek(2025, 27)
	assert.NoError(t, err)

	assert.Equal(t, "+2025-06-30", wk.StartDay().ISODate())

	wk, err = cal.FetchWeek(2024, 1)
	assert.Error(t, err)
}
